package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"wasatext/service/database"
	"wasatext/service/structs"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	var user structs.User
	err := rt.db.CheckUserUsername(userPart.Username)
	switch {
	case errors.Is(err, database.ErrNonexistentResource):
		user, err = rt.db.AddUser(userPart.Username, structs.Picture{})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	case err == nil:
		user, err = rt.db.GetUserByUsername(userPart.Username)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Prepare the response
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "Encoding error", http.StatusInternalServerError)
		return
	}
}
