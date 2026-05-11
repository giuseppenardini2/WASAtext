package api

import (
	"encoding/json"
	"net/http"
	"regexp"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	username := r.URL.Query().Get("username")
	if username == "" {
		http.Error(w, "Username is empty", http.StatusBadRequest)
	}

	usernameRegex := regexp.MustCompile("^.*?$")
	if !usernameRegex.MatchString(username) {
		http.Error(w, "Username not match regex", http.StatusBadRequest)
		return
	}

	users, err := rt.db.GetUsers(username)
	if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", "application/json")
	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, "Encoding error", http.StatusInternalServerError)
		return
	}
}
