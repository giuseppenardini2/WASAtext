package api

import (
	"encoding/json"
	"io"
	"log"
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
		log.Printf("ERRORE 30: %v\n", err.Error())
		return
	}

	if !IsUserAuthorized(userID, r, w) {
		return
	}

	conversationID, err := strconv.Atoi(ps.ByName("conversationid"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("ERRORE 31: %v\n", err.Error())
		return
	}

	if err := rt.db.CheckUserConversation(userID, conversationID); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("ERRORE 32: %v\n", err.Error())
		return
	}

	// Get request body
	uploadDirectory := "./uploads/conversations/conversation" + strconv.Itoa(conversationID) + "/" // Set up the upload directory
	if err := os.MkdirAll(uploadDirectory, os.ModePerm); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("ERRORE 33: %v\n", err.Error())
		return
	}

	err = r.ParseMultipartForm(32 << 20) // Allocate to 32 MB of memory for the image
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("ERRORE 34: %v\n", err.Error())
		return
	}

	file, fileHandler, err := r.FormFile("image") // Get the image
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("ERRORE 35: %v\n", err.Error())
		return
	}
	defer func() {
		if err := file.Close(); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Printf("ERRORE 36: %v\n", err.Error())
			return
		}
	}()

	randomGenerator := rand.New(rand.NewSource(time.Now().UnixNano()))
	fileName := time.Now().Format("20060102150405") + "-" + strconv.Itoa(userID) + "-" + strconv.Itoa(randomGenerator.Intn(100)) + filepath.Ext(fileHandler.Filename) // Set up the file name (uploadDate-userID-randomNumber.extension)

	filePath := uploadDirectory + fileName // Set up the file path
	fileDestination, err := os.Create(filePath)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("ERRORE 37: %v\n", err.Error())
		return
	}
	defer func() {
		if err := fileDestination.Close(); err != nil {
			log.Printf("ERRORE 38: %v\n", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}()

	if _, err = io.Copy(fileDestination, file); err != nil { // Copy the image file into the disk
		log.Printf("ERRORE 39: %v\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Prepare the response
	type ImagePart struct {
		ImageURI string `json:"imageURI"`
	}

	imagePart := ImagePart{
		ImageURI: filePath,
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(imagePart); err != nil {
		log.Printf("ERRORE 39.1: %v\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
