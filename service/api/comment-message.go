package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"wasatext/service/structs"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) commentMessage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Get parameters
	userID, err := strconv.Atoi(ps.ByName("userid"))
	if err != nil {
		log.Printf("ERRORE 200: %v\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !IsUserAuthorized(userID, r, w) {
		log.Printf("ERRORE 201")
		return
	}

	conversationID, err := strconv.Atoi(ps.ByName("conversationid"))
	if err != nil {
		log.Printf("ERRORE 202: %v\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := rt.db.CheckUserConversation(userID, conversationID); err != nil {
		log.Printf("ERRORE 203: %v\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	messageID, err := strconv.Atoi(ps.ByName("messageid"))
	if err != nil {
		log.Printf("ERRORE 204: %v\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := rt.db.CheckMessageConversation(messageID, conversationID); err != nil {
		log.Printf("ERRORE 205: %v\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get body request
	type ReactionPart struct {
		Emoji string `json:"emoji"`
	}
	var reactionPart ReactionPart
	if err := json.NewDecoder(r.Body).Decode(&reactionPart); err != nil {
		log.Printf("ERRORE 206: %v\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if structs.IsEmojiValid(reactionPart.Emoji) {
		log.Printf("ERRORE 207")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Prepare the response
	reaction, err := rt.db.AddReaction(userID, conversationID, messageID, reactionPart.Emoji)
	if err != nil {
		log.Printf("ERRORE 208: %v\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(reaction); err != nil {
		log.Printf("ERRORE 209: %v\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
