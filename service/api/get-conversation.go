package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"wasatext/service/structs"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getConversation(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userID, err := strconv.Atoi(ps.ByName("userID"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !IsUserAuthorized(userID, r, w) {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	recipientID, err := strconv.Atoi(r.URL.Query().Get("recipientID"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	recipientType, err := structs.GetRecipientType(r.URL.Query().Get("recipientType"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	messages, err := rt.db.GetMessagesByRecipient(userID, recipientID, recipientType)
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
