package structs

// A struct which represents a user
type User struct {
	Id       int    `json:"id"`       //User id
	Username string `json:"username"` //User username
	Picture  string `json:"picture"`  //User profile picture
}
