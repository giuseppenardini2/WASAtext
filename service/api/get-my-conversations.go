package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getMyConversations(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Get parameters
	userID, err := strconv.Atoi(ps.ByName("userid"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !IsUserAuthorized(userID, r, w) {
		return
	}

	log.Println("User authorized")

	// Prepare the response
	conversations, err := rt.db.GetConversations(userID)
	if err != nil {
		log.Printf("ERRORE: %v\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	l := len(conversations)
	log.Printf("Conversazioni lunghe %d\n", l)
	for i := range conversations {
		log.Printf("Conversazioni %d presa\n", conversations[i].ID)
	}

	log.Println("Conversazioni prese")

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(conversations); err != nil {
		log.Println("Errore JSON")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Println("Done")
}
