package structs

// A struct which represents a group
type Group struct {
	Id            int     `json:"id"`            //Group id
	Name          string  `json:"name"`          //Group name
	Picture       Picture `json:"picture"`       //Group picture
	UserCreatorID int     `json:"userCreatorID"` //Id of the user who created the group
}
