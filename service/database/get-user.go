package database

import "WASAtext/service/api/structs/User"

// Given a user ID, return the user
func (db *appdbimpl) GetUser(user_id int) (User, error) {
	var user User
	err := db.c.QueryRow("SELECT id, username, picture FROM User WHERE id=user_id").Scan(&user.id, &user.username, &user.picture)
	return user, err
}
