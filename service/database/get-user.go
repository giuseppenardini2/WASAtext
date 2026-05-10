package database

import "wasatext/service/api/structs"

// Given a user ID, return the user
func (db *appdbimpl) GetUser(userID int) (structs.User, error) {
	var user structs.User

	query := "SELECT id, username, picture FROM users WHERE id=?"
	err := db.c.QueryRow(query, userID).Scan(&user.ID, &user.Username, &user.Picture)
	return user, err
}
