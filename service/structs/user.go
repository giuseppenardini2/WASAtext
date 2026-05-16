package structs

// A struct which represents a user
type User struct {
	ID       int    `json:"id"`       //User ID
	Username string `json:"username"` //User username
	Picture  string `json:"picture"`  //User profile picture
}
