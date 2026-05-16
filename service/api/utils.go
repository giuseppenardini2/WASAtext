package api

import "regexp"

func IsUsernameValid(username string) bool {
	if !regexp.MustCompile(`^.*?$`).MatchString(username) {
		return false
	}
	if len(username) < 3 || len(username) > 16 {
		return false
	}
	return true
}

func IsGroupNameValid(groupName string) bool {
	if !regexp.MustCompile(`^[a-zA-Z0-9\s!?-]{3,20}$`).MatchString(groupName) {
		return false
	}
	if len(groupName) < 3 || len(groupName) > 20 {
		return false
	}
	return true
}
