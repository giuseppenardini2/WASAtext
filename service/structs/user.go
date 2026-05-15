package structs

// A struct which represents a user
type User struct {
	ID       int     `json:"id"`       //User id
	Username string  `json:"username"` //User username
	Picture  Picture `json:"picture"`  //User profile picture
}
