package api

import (
	"net/http"
	"regexp"
	"wasatext/service/structs"

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

	var user structs.User
	if is_username_exists == false {
		//Creo lo user
	}

	//Prendo lo user

	/*username := r.URL.Query().Get("username")
	if username == "" {
		http.Error(w, "Username is empty", http.StatusBadRequest)
		return
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
	}*/
}
