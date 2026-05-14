package api

import "regexp"

func IsUsernameRegexValid(username string) bool {
	return regexp.MustCompile("^.*?$").MatchString(username)
}
