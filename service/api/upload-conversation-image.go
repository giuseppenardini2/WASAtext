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

func (rt *_router) uploadConversationImage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Get parameters
	userID, err := strconv.Atoi(ps.ByName("userid"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !IsUserAuthorized(userID, r, w) {
		return
	}

	recipientID, err := strconv.Atoi(r.URL.Query().Get("recipientid"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := rt.db.CheckUserConversation(userID, recipientID); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get request body
	currentWorkDirectory := "./"
	uploadDirectory := "uploads/conversations/conversation" + strconv.Itoa(userID) + "/" // Set up the upload directory
	if err := os.MkdirAll(currentWorkDirectory+uploadDirectory, os.ModePerm); err != nil {
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

	filePath := currentWorkDirectory + uploadDirectory + fileName // Set up the file path
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
	type ImagePart struct {
		ImageURI string `json:"imageURI"`
	}

	imagePart := ImagePart{
		ImageURI: uploadDirectory + fileName,
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(imagePart); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
