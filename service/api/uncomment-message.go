package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) uncommentMessage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Get parameters
	userID, err := strconv.Atoi(ps.ByName("userid"))
	if err != nil {
		log.Printf("ERRORE 300: %v\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !IsUserAuthorized(userID, r, w) {
		log.Printf("ERRORE 301\n")
		return
	}

	conversationID, err := strconv.Atoi(ps.ByName("conversationid"))
	if err != nil {
		log.Printf("ERRORE 302: %v\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := rt.db.CheckUserConversation(userID, conversationID); err != nil {
		log.Printf("ERRORE 303: %v\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	messageID, err := strconv.Atoi(ps.ByName("messageid"))
	if err != nil {
		log.Printf("ERRORE 304: %v\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	reactionID, err := strconv.Atoi(ps.ByName("reactionid"))
	if err != nil {
		log.Printf("ERRORE 305: %v\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := rt.db.CheckUserReaction(userID, reactionID, messageID, conversationID); err != nil {
		log.Printf("ERRORE 306: %v\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Prepare the response
	if err := rt.db.DeleteReaction(reactionID, conversationID, messageID); err != nil {
		log.Printf("ERRORE 307: %v\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
