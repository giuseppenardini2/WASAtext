package database

import "wasatext/service/api/structs"

// Given a username, return the users that corresponds to that username
func (db *appdbimpl) GetUsers(username string) (structs.User[], error) {
	raws, err := db.c.Query("SELECT id, username, picture FROM User WHERE username LIKE concat(?, '%')", username)
	if err != nil {
		return nil, err
	}

	var users structs.Users[]
	for rows.Next(){
		var user structs.User
		if err := rows.Scan(&user.id, &user.username, &user.picture); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return users, err
}
