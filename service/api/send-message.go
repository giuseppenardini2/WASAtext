package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) sendMessage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	recipientID, err := strconv.Atoi(r.URL.Query().Get("recipientID"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	messages, err := rt.db.GetMessages(recipientID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", "application/json")
	if err := json.NewEncoder(w).Encode(messages); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
