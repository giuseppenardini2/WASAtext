package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"wasatext/service/database"
	"wasatext/service/structs"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) createConversation(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	// Get body request
	type ConversationPart struct {
		SecondUserID int `json:"secondUserID"`
	}
	var conversationPart ConversationPart
	if err := json.NewDecoder(r.Body).Decode(&conversationPart); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := rt.db.CheckUserPrivateConversation(userID, conversationPart.SecondUserID); err != database.NonexistentResource {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Prepare the response
	conversationID, err := rt.db.AddPrivateConversation(userID, conversationPart.SecondUserID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	conversation := structs.Conversation{
		ID:       int(conversationID),
		Messages: nil,
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("content-type", "application/json")
	if err := json.NewEncoder(w).Encode(conversation); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
