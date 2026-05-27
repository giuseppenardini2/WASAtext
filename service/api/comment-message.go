package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"wasatext/service/structs"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) commentMessage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Get parameters
	userID, err := strconv.Atoi(ps.ByName("userID"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !IsUserAuthorized(userID, r, w) {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	conversationID, err := strconv.Atoi(ps.ByName("conversationID"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := rt.db.CheckUserConversation(userID, conversationID); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	messageID, err := strconv.Atoi(ps.ByName("messageID"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := rt.db.CheckMessageConversation(messageID, conversationID); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get body request
	type ReactionPart struct {
		Emoji string `json:"emoji"`
	}
	var reactionPart ReactionPart
	if err := json.NewDecoder(r.Body).Decode(&reactionPart); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if structs.IsEmojiValid(reactionPart.Emoji) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Prepare the response
	if err := rt.db.AddReaction(userID, conversationID, messageID, reactionPart.Emoji); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	reaction, err := rt.db.GetReaction(userID, conversationID, messageID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("content-type", "application/json")
	if err := json.NewEncoder(w).Encode(reaction); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
