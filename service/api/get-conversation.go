package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getConversation(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Get parameters
	userID, err := strconv.Atoi(ps.ByName("userid"))
	if err != nil {
		log.Printf("ERRORE 10: %v\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !IsUserAuthorized(userID, r, w) {
		return
	}

	recipientID, err := strconv.Atoi(r.URL.Query().Get("recipientid"))
	if err != nil {
		log.Printf("ERRORE 11: %v\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := rt.db.CheckUserConversation(userID, recipientID); err != nil {
		log.Printf("ERRORE 12: %v\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Prepare the response
	conversation, err := rt.db.GetConversation(userID, recipientID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("ERRORE 13: %v\n", err.Error())
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(conversation); err != nil {
		log.Printf("ERRORE 14: %v\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
