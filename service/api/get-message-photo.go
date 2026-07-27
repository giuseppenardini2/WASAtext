package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getMessagePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Get parameters
	userID, err := strconv.Atoi(ps.ByName("userid"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("ERRORE 50: %v\n", err.Error())
		return
	}

	if !IsUserAuthorized(userID, r, w) {
		return
	}

	conversationID, err := strconv.Atoi(ps.ByName("conversationid"))
	if err != nil {
		log.Printf("ERRORE 51: %v\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := rt.db.CheckUserConversation(userID, conversationID); err != nil {
		log.Printf("ERRORE 52: %v\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	messageID, err := strconv.Atoi(ps.ByName("messageid"))
	if err != nil {
		log.Printf("ERRORE 53: %v\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := rt.db.CheckMessageConversation(messageID, conversationID); err != nil {
		log.Printf("ERRORE 54: %v\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Prepare the response
	pictureURI, err := rt.db.GetMessagePictureURI(conversationID, messageID)
	if err != nil {
		log.Printf("ERRORE 55: %v\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if pictureURI == "" {
		w.WriteHeader(http.StatusNotFound)
		log.Printf("ERRORE 56")
		return
	}

	http.ServeFile(w, r, pictureURI)
}
