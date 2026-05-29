package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"wasatext/service/structs"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) createGroup(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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
	type GroupPart struct {
		Name         string         `json:"name"`
		Partecipants []structs.User `json:"partecipants"`
	}
	var groupPart GroupPart
	if err := json.NewDecoder(r.Body).Decode(&groupPart); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for i := range groupPart.Partecipants {
		if err := rt.db.CheckUser(groupPart.Partecipants[i].ID); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	// Prepare the response
	groupConversation, err := rt.db.AddGroupConversation(groupPart.Name, "", userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	for i := range groupPart.Partecipants {
		if err := rt.db.AddUserToGroupConversation(groupPart.Partecipants[i].ID, groupConversation.ID); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("content-type", "application/json")
	if err := json.NewEncoder(w).Encode(groupConversation); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
