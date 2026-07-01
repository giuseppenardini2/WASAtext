package api

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"wasatext/service/database"
	"wasatext/service/structs"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) createConversation(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Get parameters
	userID, err := strconv.Atoi(ps.ByName("userid"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !IsUserAuthorized(userID, r, w) {
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

	if err := rt.db.CheckUserPrivateConversation(userID, conversationPart.SecondUserID); !errors.Is(err, database.ErrNonexistentResource) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Prepare the response
	conversationID, err := rt.db.AddPrivateConversation(userID, conversationPart.SecondUserID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("Conversazione creata: %d\n", conversationID)

	conversation := structs.Conversation{
		ID:       int(conversationID),
		Messages: nil,
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(conversation); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
