package api

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
)

// Given an HTTP request, extract the bearer token from the Authorization header
func getUserToken(r *http.Request) (int, error) {
	// Get the authorization header
	authorizationHeader := r.Header.Get("Authorization")
	if authorizationHeader == "" {
		return -1, errors.New("Authorization header is empty")
	}

	// Split the line to get the bearer token (format: "Bearer token123")
	authorizationParts := strings.Split(authorizationHeader, " ")
	if len(authorizationParts) != 2 || !strings.EqualFold(authorizationParts[0], "Bearer") {
		return -1, errors.New("Authorization header error")
	}

	// Convert the bearer token (which is the user ID) into integer
	token, err := strconv.Atoi(authorizationParts[1])
	if err != nil {
		return -1, err
	}
	return token, nil
}

// Given a user ID and the HTTP request, verify if the user corresponds to the token in the authorization header
func IsUserAuthorized(userID int, r *http.Request, w http.ResponseWriter) bool {
	token, err := getUserToken(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return false
	}
	if userID != token {
		w.WriteHeader(http.StatusForbidden)
		return false
	}
	return true
}
