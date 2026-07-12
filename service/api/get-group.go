package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getGroup(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Get parameters
	userID, err := strconv.Atoi(ps.ByName("userid"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !IsUserAuthorized(userID, r, w) {
		return
	}

	peerID, err := strconv.Atoi(r.URL.Query().Get("peerid"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := rt.db.CheckGroupConversation(peerID); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := rt.db.CheckGroupConversationMember(userID, peerID); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Prepare the response
	groupConversation, err := rt.db.GetGroupConversation(peerID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(groupConversation); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
