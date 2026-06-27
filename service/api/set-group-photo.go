package api

import (
	"encoding/json"
	"io"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

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
	uploadDirectory := "./uploads/groups/group" + strconv.Itoa(groupID) + "/" // Set up the upload directory
	if err := os.MkdirAll(uploadDirectory, os.ModePerm); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = r.ParseMultipartForm(32 << 20) // Allocate to 32 MB of memory for the image
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	file, fileHandler, err := r.FormFile("image") // Get the image
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer func() {
		if err := file.Close(); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}()

	randomGenerator := rand.New(rand.NewSource(time.Now().UnixNano()))
	fileName := time.Now().Format("20060102150405") + "-" + strconv.Itoa(userID) + "-" + strconv.Itoa(randomGenerator.Intn(100)) + filepath.Ext(fileHandler.Filename) // Set up the file name (uploadDate-userID-randomNumber.extension)

	filePath := uploadDirectory + fileName // Set up the file path
	fileDestination, err := os.Create(filePath)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer func() {
		if err := fileDestination.Close(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}()

	if _, err = io.Copy(fileDestination, file); err != nil { // Copy the image file into the disk
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Prepare the response
	if err := rt.db.UpdateGroupConversationPicture(groupID, filePath); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	type ImagePart struct {
		ImageURI string `json:"imageURI"`
	}

	imagePart := ImagePart{
		ImageURI: filePath,
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(imagePart); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
