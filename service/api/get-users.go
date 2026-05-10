package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// getHelloWorld is an example of HTTP endpoint that returns "Hello world!" as a plain text
func (rt *_router) getUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	/*w.Header().Set("content-type", "text/plain")
	_, _ = w.Write([]byte("Hello World!"))*/

	username := r.URL.Query().Get("username")

	users, err := rt.db.GetUsers(username)
	if err != nil {
		return
	}

	w.Header().Set("content-type", "application/json")
	if err := json.NewEncoder(w).Encode(users); err != nil {
		//Indicare il tipo di errore, return
		return
	}
}
