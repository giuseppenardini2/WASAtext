package structs

import "regexp"

type User struct {
	ID       int     `json:"id"`       // User id
	Username string  `json:"username"` // User username
	Picture  Picture `json:"picture"`  // User profile picture
}

func IsUsernameValid(username string) bool {
	if username == "" || (len(username) < 3 || len(username) > 16) {
		return false
	}
	if !regexp.MustCompile(`^.*?$`).MatchString(username) {
		return false
	}
	return true
}
