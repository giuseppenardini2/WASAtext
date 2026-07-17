package api

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"wasatext/service/database"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) addToGroup(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Get parameters
	userID, err := strconv.Atoi(ps.ByName("userid"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("ERRORE 0: %v\n", err.Error())
		return
	}

	if !IsUserAuthorized(userID, r, w) {
		log.Printf("Non autorizzato\n")
		return
	}

	groupID, err := strconv.Atoi(ps.ByName("groupid"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("ERRORE 1: %v\n", err.Error())
		return
	}

	if err := rt.db.CheckGroupConversation(groupID); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("ERRORE 2: %v\n", err.Error())
		return
	}

	if err := rt.db.CheckGroupConversationMember(userID, groupID); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("ERRORE 3: %v\n", err.Error())
		return
	}

	userToAddID, err := strconv.Atoi(r.URL.Query().Get("peerid"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("ERRORE 4: %v\n", err.Error())
		return
	}

	if err := rt.db.CheckUser(userToAddID); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("ERRORE 5: %v\n", err.Error())
		return
	}

	if err := rt.db.CheckGroupConversationMember(userToAddID, groupID); !errors.Is(err, database.ErrNonexistentResource) {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("ERRORE 6: %v\n", err.Error())
		return
	}

	// Prepare the response
	if err := rt.db.AddGroupConversationMember(userToAddID, groupID); err != nil {
		log.Printf("ERRORE 7: %v\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	type Response struct {
		UserID int `json:"id"`
	}
	response := Response{
		UserID: userToAddID,
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("ERRORE 8: %v\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
