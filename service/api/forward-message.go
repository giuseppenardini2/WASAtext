package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) forwardMessage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Get parameters
	userID, err := strconv.Atoi(ps.ByName("userid"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !IsUserAuthorized(userID, r, w) {
		return
	}

	conversationID, err := strconv.Atoi(ps.ByName("conversationid"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := rt.db.CheckUserConversation(userID, conversationID); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	messageID, err := strconv.Atoi(ps.ByName("messageid"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	recipientID, err := strconv.Atoi(r.URL.Query().Get("recipientid"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := rt.db.CheckUserConversation(userID, recipientID); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get body request
	type MessagePart struct {
		Body       string `json:"body"`
		PictureURI string `json:"pictureURI"`
	}
	var messagePart MessagePart
	if err := json.NewDecoder(r.Body).Decode(&messagePart); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if messagePart.Body == "" && messagePart.PictureURI == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Prepare the response
	message, err := rt.db.AddForwardedMessage(messageID, conversationID, userID, recipientID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(message); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
