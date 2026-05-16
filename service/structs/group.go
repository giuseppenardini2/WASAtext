package structs

// A struct which represents a group
type Group struct {
	Id            int    `json:"id"`            //Group ID
	Name          string `json:"name"`          //Group name
	Picture       string `json:"picture"`       //Group picture
	UserCreatorID int    `json:"userCreatorID"` //ID of the user who created the group
}
