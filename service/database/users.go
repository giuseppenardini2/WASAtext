package database

import (
	"wasatext/service/structs"
)

/*
 * ADD OPERATIONS
 */

func (db *appdbimpl) AddUser(username string, picture string) (structs.User, error) {
	query := "INSERT INTO users (username) VALUES (?)"
	result, err := db.c.Exec(query, username)
	if err != nil {
		return structs.User{}, err
	}

	userID, err := result.LastInsertId()
	if err != nil {
		return structs.User{}, err
	}

	if picture != "" {
		query := "UPDATE users SET picture=? WHERE id=?"
		if _, err := db.c.Exec(query, picture, userID); err != nil {
			return structs.User{}, err
		}
	}

	return structs.User{
		ID:       int(userID),
		Username: username,
		Picture:  picture,
	}, nil
}

/*
 * UPDATE OPERATIONS
 */

func (db *appdbimpl) UpdateUserUsername(userID int, username string) error {
	query := "UPDATE users SET username=? WHERE id=?"
	_, err := db.c.Exec(query, username, userID)
	return err
}

func (db *appdbimpl) UpdateUserPicture(userID int, picture string) error {
	query := "UPDATE users SET picture=? WHERE id=?"
	_, err := db.c.Exec(query, picture, userID)
	return err
}

/*
 * GET OPERATIONS
 */

func (db *appdbimpl) GetUser(userID int) (structs.User, error) {
	var user structs.User

	query := "SELECT id, username, picture FROM users WHERE id=?"
	err := db.c.QueryRow(query, userID).Scan(&user.ID, &user.Username, &user.Picture)
	return user, err
}

func (db *appdbimpl) GetUserByUsername(username string) (structs.User, error) {
	var user structs.User

	query := "SELECT id, username, picture FROM users WHERE username=?"
	err := db.c.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Picture)
	return user, err
}

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

func (db *appdbimpl) CheckUser(userID int) error {
	var count int

	query := "SELECT COUNT(id) FROM users WHERE id=?"
	err := db.c.QueryRow(query, userID).Scan(&count)
	if err != nil {
		return err
	}
	if count == 0 {
		return ErrNonexistentResource
	}
	return nil
}

func (db *appdbimpl) CheckUserUsername(username string) error {
	var count int

	query := "SELECT COUNT(id) FROM users WHERE username=?"
	err := db.c.QueryRow(query, username).Scan(&count)
	if err != nil {
		return err
	}
	if count == 0 {
		return ErrNonexistentResource
	}
	return nil
}
