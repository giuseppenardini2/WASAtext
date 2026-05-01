package database

import "wasatext/service/api/structs"

// Given a user ID, return the user
func (db *appdbimpl) GetUser(id int) (structs.User, error) {
	var user structs.User

	var query = "SELECT id, username, picture FROM users WHERE id=?"
	err := db.c.QueryRow(query, id).Scan(&user.Id, &user.Username, &user.Picture)
	return user, err
}
