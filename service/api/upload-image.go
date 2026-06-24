package api

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) uploadImage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Get parameters
	userID, err := strconv.Atoi(ps.ByName("userid"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !IsUserAuthorized(userID, r, w) {
		return
	}

	// Get body request
	uploadDirectory := "./uploads/user" + strconv.Itoa(userID) + "/images/" // Set up the upload directory
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

	fileName := time.Now().String() + fileHandler.Filename

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
	type ImagePart struct {
		ImageURI string `json:"pictureURI"`
	}

	picturePart := ImagePart{ // Set the image URI
		ImageURI: "http://" + r.Host + uploadDirectory + fileName,
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(picturePart); err != nil {
		http.Error(w, "Encoding error", http.StatusInternalServerError)
		return
	}
}
