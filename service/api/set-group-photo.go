package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) setGroupPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := rt.db.CheckGroupConversationMember(userID, groupID); err != nil {
		w.WriteHeader(http.StatusBadRequest)
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
	if err := rt.db.UpdateGroupConversationPicture(groupID, picturePart.PictureURI); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("The request succeeded"))
}
