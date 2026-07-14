package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getGroup(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Get parameters
	userID, err := strconv.Atoi(ps.ByName("userid"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("ERRORE 0: %v\n", err.Error())
		return
	}

	if !IsUserAuthorized(userID, r, w) {
		return
	}

	peerID, err := strconv.Atoi(r.URL.Query().Get("peerid"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("ERRORE 1: %v\n", err.Error())
		return
	}

	if err := rt.db.CheckGroupConversation(peerID); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("ERRORE 2: %v\n", err.Error())
		return
	}

	if err := rt.db.CheckGroupConversationMember(userID, peerID); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("ERRORE 3: %v\n", err.Error())
		return
	}

	// Prepare the response
	groupConversation, err := rt.db.GetGroupConversation(peerID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("ERRORE 4: %v\n", err.Error())
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(groupConversation); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
