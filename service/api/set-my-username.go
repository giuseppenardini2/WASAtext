package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"wasatext/service/structs"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) setMyUserName(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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
	type UserPart struct {
		Username string `json:"username"`
	}
	var userPart UserPart
	if err := json.NewDecoder(r.Body).Decode(&userPart); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !structs.IsUsernameValid(userPart.Username) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := rt.db.CheckUserUsername(userPart.Username); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Prepare the response
	if err := rt.db.UpdateUserUsername(userID, userPart.Username); err != nil {
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
