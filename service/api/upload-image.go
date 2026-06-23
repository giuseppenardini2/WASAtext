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
	uploadDirectory := "./uploads/" + strconv.Itoa(123) + "/images/" // Set up the folder
	if err := os.MkdirAll(uploadDirectory, os.ModePerm); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = r.ParseMultipartForm(32 << 20) // Allocate to 32 MB of memory for the file
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile("image") // Get the image
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer func() {
		err = file.Close()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}()

	fileName := time.Now().String() // Set up the file name

	filePath := uploadDirectory + fileName // Set up the file path
	fileDestination, err := os.Create(filePath)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer func() {
		err = fileDestination.Close()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}()

	if _, err = io.Copy(fileDestination, file); err != nil { // Copy the file
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Prepare the response
	type PicturePart struct {
		PictureURI string `json:"pictureURI"`
	}

	picturePart := PicturePart{ // Set the image URI
		PictureURI: filePath,
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(picturePart); err != nil {
		http.Error(w, "Encoding error", http.StatusInternalServerError)
		return
	}
}
