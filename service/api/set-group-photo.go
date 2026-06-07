package api

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"wasatext/service/structs"

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
		w.WriteHeader(http.StatusForbidden)
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
	file, _, err := r.FormFile("image")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	data, err := io.ReadAll(file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fileType := http.DetectContentType(data)
	if fileType != "image/jpeg" && fileType != "image/png" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Prepare the response
	picture := structs.Picture{
		Data: data,
	}

	if err := rt.db.UpdateGroupConversationPicture(groupID, picture); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", "application/json")
	if err := json.NewEncoder(w).Encode(picture); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
