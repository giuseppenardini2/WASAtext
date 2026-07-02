package api

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getMyPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Get parameters
	userID, err := strconv.Atoi(ps.ByName("userid"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !IsUserAuthorized(userID, r, w) {
		return
	}

	peerID, err := strconv.Atoi(r.URL.Query().Get("peerid"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := rt.db.CheckUser(peerID); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Prepare the response
	pictureURI, err := rt.db.GetUserPictureURI(peerID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if pictureURI == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	http.ServeFile(w, r, pictureURI)
}
