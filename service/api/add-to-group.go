package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"wasatext/service/database"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) addToGroup(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	userToAddID, err := strconv.Atoi(r.URL.Query().Get("userToAddID"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := rt.db.CheckUser(userToAddID); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := rt.db.CheckGroupConversationMember(userToAddID, groupID); err != database.ErrNonexistentResource {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Prepare the response
	if err := rt.db.AddGroupConversationMember(userID, groupID); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	type Response struct {
		UserID int `json:"id"`
	}
	response := Response{
		UserID: userToAddID,
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("content-type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
