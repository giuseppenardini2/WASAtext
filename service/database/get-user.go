package database

import "WASAtext/service/api/structs"

// Given a user ID, return the user
func (db *appdbimpl) GetUser(user_id int) (structs.User, error) {
	var user structs.User
	err := db.c.QueryRow("SELECT id, username, picture FROM User WHERE id=?", user_id).Scan(&user.id, &user.username, &user.picture)
	return user, err
}
