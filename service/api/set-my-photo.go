package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) setMyPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Get parameters
	userID, err := strconv.Atoi(ps.ByName("userid"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !IsUserAuthorized(userID, r, w) {
		return
	}

	// Get request body
	type PicturePart struct {
		PictureURI string `json:"pictureURI"`
	}
	var picturePart PicturePart
	if err := json.NewDecoder(r.Body).Decode(&picturePart); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Prepare the response
	if err := rt.db.UpdateUserPicture(userID, picturePart.PictureURI); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte("The request succeeded")); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
