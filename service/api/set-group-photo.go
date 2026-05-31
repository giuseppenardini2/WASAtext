package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) setGroupPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	groupID, err := strconv.Atoi(ps.ByName("groupID"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := rt.db.CheckGroupConversationMember(userID, groupID); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get request body
	type GroupPart struct {
		Picture string `json:"picture"`
	}
	var userPart GroupPart
	if err := json.NewDecoder(r.Body).Decode(&userPart); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// TODO: Aggiungere controlli foto

	// Prepare the response
	if err := rt.db.UpdateGroupConversationPicture(groupID, userPart.Picture); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", "application/json")
	if err := json.NewEncoder(w).Encode(userPart); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
