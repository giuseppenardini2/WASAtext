package structs

import "regexp"

// A struct which represents a group
type Group struct {
	Id            int    `json:"id"`            //Group ID
	Name          string `json:"name"`          //Group name
	Picture       string `json:"picture"`       //Group picture
	UserCreatorID int    `json:"userCreatorID"` //ID of the user who created the group
}

// Check if a group name is valid
func IsGroupNameValid(groupName string) bool {
	if groupName == "" || (len(groupName) < 3 || len(groupName) > 20) {
		return false
	}
	if !regexp.MustCompile(`^[a-zA-Z0-9\s!?-]{3,20}$`).MatchString(groupName) {
		return false
	}
	return true
}
