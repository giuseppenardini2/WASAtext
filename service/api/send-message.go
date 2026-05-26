package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) sendMessage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	replyMessageID := -1
	if r.URL.Query().Has("replyMessageID") {
		var err error
		replyMessageID, err = strconv.Atoi(r.URL.Query().Get("replyMessageID"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	// Get body request
	type MessagePart struct {
		Body    string `json:"body"`
		Picture string `json:"picture"`
	}
	var messagePart MessagePart
	if err := json.NewDecoder(r.Body).Decode(&messagePart); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if messagePart.Body == "" && messagePart.Picture == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Prepare the response
	var messageID int64 // Add the message to DB
	if replyMessageID != -1 {
		var err error
		messageID, err = rt.db.AddReplyMessage(messagePart.Body, messagePart.Picture, userID, conversationID, replyMessageID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	} else {
		var err error
		messageID, err = rt.db.AddMessage(messagePart.Body, messagePart.Picture, userID, conversationID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	message, err := rt.db.GetMessage(conversationID, int(messageID)) // TODO: Rivedere. Tutto int64 o no?
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("content-type", "application/json")
	if err := json.NewEncoder(w).Encode(message); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
