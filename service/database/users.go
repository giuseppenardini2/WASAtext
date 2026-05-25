package database

import "wasatext/service/structs"

/*
 * ADD OPERATIONS
 */
// Add a user and return its ID
func (db *appdbimpl) AddUser(username string) (int64, error) {
	query := "INSERT INTO users (username) VALUES (?)"
	result, err := db.c.Exec(query, username)
	if err != nil {
		return -1, err
	}
	return result.LastInsertId()
}

/*
 * UPDATE OPERATIONS
 */
// Given a user ID and a username, update the username of the user
func (db *appdbimpl) UpdateUserUsername(userID int, username string) error {
	query := "UPDATE users SET username=? WHERE id=?"
	_, err := db.c.Exec(query, username, userID)
	return err
}

// Given a user ID and a profile picture, update the profile picture of the user
func (db *appdbimpl) UpdateUserPicture(userID int, picture string) error {
	query := "UPDATE users SET picture=? WHERE id=?"
	_, err := db.c.Exec(query, picture, userID)
	return err
}

/*
 * GET OPERATIONS
 */
// Given a user ID, return the user
func (db *appdbimpl) GetUser(userID int) (structs.User, error) {
	var user structs.User

	query := "SELECT id, username, picture FROM users WHERE id=?"
	err := db.c.QueryRow(query, userID).Scan(&user.ID, &user.Username, &user.Picture)
	return user, err
}

// Given a username, return the user
func (db *appdbimpl) GetUserByUsername(username string) (structs.User, error) {
	var user structs.User

	query := "SELECT id, username, picture FROM users WHERE username=?"
	err := db.c.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Picture)
	return user, err
}

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

/*
 * CHECK OPERATIONS
 */
// Given a username, return if the username is already used
func (db *appdbimpl) CheckUserUsername(username string) (bool, error) {
	var count int

	query := "SELECT COUNT(id) FROM users WHERE username=?"
	err := db.c.QueryRow(query, username).Scan(&count)
	if err != nil {
		return false, err
	}
	return count != 0, err
}
