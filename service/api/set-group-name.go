package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"wasatext/service/database"
	"wasatext/service/structs"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) setGroupName(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	err = rt.db.CheckGroupConversation(groupID)
	if errors.Is(err, database.ErrNonexistentResource) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := rt.db.CheckGroupConversationMember(userID, groupID); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get body request
	type GroupPart struct {
		Name string `json:"name"`
	}
	var groupPart GroupPart
	if err := json.NewDecoder(r.Body).Decode(&groupPart); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !structs.IsGroupNameValid(groupPart.Name) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Prepare the response
	if err := rt.db.UpdateGroupConversationName(userID, groupPart.Name); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", "application/json")
	if err := json.NewEncoder(w).Encode(groupPart); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
