package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) setMyUserName(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userID, err := strconv.Atoi(ps.ByName("userID"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !IsUserAuthorized(userID, r, w) {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	username := r.URL.Query().Get("username")
	if !IsUsernameValid(username) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	is_username_used, err := rt.db.CheckUserUsername(username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if is_username_used {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := rt.db.UpdateUserUsername(userID, username); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", "application/json")
	if err := json.NewEncoder(w).Encode(username); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
