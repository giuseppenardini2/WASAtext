package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getReactions(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Get parameters
	userID, err := strconv.Atoi(ps.ByName("userid"))
	if err != nil {
		log.Printf("ERRORE 100: %v\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !IsUserAuthorized(userID, r, w) {
		log.Printf("ERRORE 101")
		return
	}

	conversationID, err := strconv.Atoi(ps.ByName("conversationid"))
	if err != nil {
		log.Printf("ERRORE 102: %v\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := rt.db.CheckUserConversation(userID, conversationID); err != nil {
		log.Printf("ERRORE 103: %v\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	messageID, err := strconv.Atoi(ps.ByName("messageid"))
	if err != nil {
		log.Printf("ERRORE 104: %v\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Prepare the response
	reactions, err := rt.db.GetReactions(conversationID, messageID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("ERRORE 105: %v\n", err.Error())
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(reactions); err != nil {
		log.Printf("ERRORE 106: %v\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
