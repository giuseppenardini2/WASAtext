package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"wasatext/service/structs"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) sendMessage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Get parameters
	userID, err := strconv.Atoi(ps.ByName("userid"))
	if err != nil {
		log.Printf("ERRORE 0: %v\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !IsUserAuthorized(userID, r, w) {
		log.Printf("ERRORE 1")
		return
	}

	conversationID, err := strconv.Atoi(ps.ByName("conversationid"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("ERRORE 2: %v\n", err.Error())
		return
	}

	if err := rt.db.CheckUserConversation(userID, conversationID); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("ERRORE 3: %v\n", err.Error())
		return
	}

	replyMessageID := -1
	if r.URL.Query().Has("replymessageid") {
		var err error
		replyMessageID, err = strconv.Atoi(r.URL.Query().Get("replymessageid"))
		if err != nil {
			log.Printf("ERRORE 4: %v\n", err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	// Get body request
	type MessagePart struct {
		Body       string `json:"body"`
		PictureURI string `json:"pictureURI"`
	}
	var messagePart MessagePart
	if err := json.NewDecoder(r.Body).Decode(&messagePart); err != nil {
		log.Printf("ERRORE 5: %v\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if messagePart.Body == "" && messagePart.PictureURI == "" {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("ERRORE 6\n")
		return
	}

	// Prepare the response
	var message structs.Message
	if replyMessageID != -1 {
		var err error
		message, err = rt.db.AddReplyMessage(messagePart.Body, messagePart.PictureURI, userID, conversationID, replyMessageID)
		if err != nil {
			log.Printf("ERRORE 7: %v\n", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	} else {
		var err error
		message, err = rt.db.AddMessage(messagePart.Body, messagePart.PictureURI, userID, conversationID)
		if err != nil {
			log.Printf("ERRORE 8: %v\n", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(message); err != nil {
		log.Printf("ERRORE 9: %v\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
