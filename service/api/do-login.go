package api

import (
	"encoding/json"
	"net/http"
	"wasatext/service/structs"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Get parameters
	username := ps.ByName("username")
	if username == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !structs.IsUsernameValid(username) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := rt.db.CheckUserUsername(username); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Prepare the response
	user, err := rt.db.AddUser(username, "")
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
