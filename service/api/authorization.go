package api

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
)

// Given an HTTP request, extract the bearer token from the Authorization header
func (rt *_router) GetUserToken(r *http.Request) (int, error) {
	// Get the authorization header
	authorizationHeader := r.Header.Get("Authorization")
	if authorizationHeader == "" {
		return -1, errors.New("Authorization header error")
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

// TODO: La funzione verifica se l'utente è quello del token
func IsUserAuthorized(userID int, token int) bool {
	return true
}
