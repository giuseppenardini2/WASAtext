package api

import (
	"encoding/json"
	"net/http"
	"regexp"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	username := ps.ByName("username")
	if username == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	usernameRegex := regexp.MustCompile("^.*?$")
	if !usernameRegex.MatchString(username) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	is_username_exists, err := rt.db.CheckUserUsername(username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if is_username_exists == false {
		_, err := rt.db.AddUser(username, nil) //Come rendere nulla la stringa ?
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	user, err := rt.db.GetUserByUsername(username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "Encoding error", http.StatusInternalServerError)
		return
	}
}
