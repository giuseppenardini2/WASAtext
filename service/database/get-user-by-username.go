package database

import "wasatext/service/api/structs"

// Given a username, return the user
func (db *appdbimpl) GetUserByUsername(username string) (structs.User, error) {
	var user structs.User

	query := "SELECT id, username, picture FROM users WHERE username=?"
	err := db.c.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Picture)
	return user, err
}
