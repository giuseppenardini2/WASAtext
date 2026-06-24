package structs

import "regexp"

var usernameRegex = regexp.MustCompile(`^[A-Za-z0-9._]{3,16}$`)

type User struct {
	ID         int    `json:"id"`         // User id
	Username   string `json:"username"`   // User username
	PictureURI string `json:"pictureURI"` // User profile picture
}

func IsUsernameValid(username string) bool {
	if username == "" || (len(username) < 3 || len(username) > 16) {
		return false
	}
	if !usernameRegex.MatchString(username) {
		return false
	}
	return true
}
