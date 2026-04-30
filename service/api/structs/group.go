package structs

// A struct which represents a group
type Group struct {
	Id          int    `json:"id"`          //Group id
	Name        string `json:"name"`        //Group name
	Picture     string `json:"picture"`     //Group picture
	UserCreator int    `json:"userCreator"` //Id of the user who created the group
}
