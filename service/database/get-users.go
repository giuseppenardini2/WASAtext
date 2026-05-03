package database

import "wasatext/service/api/structs"

// Given a username, return the users that corresponds to that username
func (db *appdbimpl) GetUsers(username string) ([]structs.User, error) {
	query := "SELECT id, username, picture FROM users WHERE username LIKE concat(?, '%')"
	rows, err := db.c.Query(query, username)
	if err != nil {
		return nil, err
	}

	var users []structs.User
	for rows.Next() {
		var user structs.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Picture); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return users, err
}
