package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getMyConversations(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Get parameters
	userID, err := strconv.Atoi(ps.ByName("userid"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !IsUserAuthorized(userID, r, w) {
		return
	}

	// Prepare the response
	conversations, err := rt.db.GetConversations(userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(conversations); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
