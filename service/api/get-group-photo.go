package api

import (
	"errors"
	"net/http"
	"strconv"
	"wasatext/service/database"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getGroupPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Get parameters
	userID, err := strconv.Atoi(ps.ByName("userid"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !IsUserAuthorized(userID, r, w) {
		return
	}

	groupID, err := strconv.Atoi(ps.ByName("groupid"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = rt.db.CheckGroupConversation(groupID)
	if errors.Is(err, database.ErrNonexistentResource) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := rt.db.CheckGroupConversationMember(userID, groupID); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Prepare the response
	pictureURI, err := rt.db.GetGroupConversationPictureURI(groupID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if pictureURI == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	http.ServeFile(w, r, "./"+pictureURI)
}
