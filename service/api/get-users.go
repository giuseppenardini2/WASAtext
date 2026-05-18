package api

import (
	"encoding/json"
	"net/http"
	"wasatext/service/structs"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	username := r.URL.Query().Get("username")
	if !structs.IsUsernameValid(username) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	users, err := rt.db.GetUsers(username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", "application/json")
	if err := json.NewEncoder(w).Encode(users); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
