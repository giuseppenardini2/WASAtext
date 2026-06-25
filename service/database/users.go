package database

import (
	"wasatext/service/structs"
)

/*
 * ADD OPERATIONS
 */

func (db *appdbimpl) AddUser(username string, pictureURI string) (structs.User, error) {
	query := "INSERT INTO users (username) VALUES (?)"
	result, err := db.c.Exec(query, username)
	if err != nil {
		return structs.User{}, err
	}

	userID, err := result.LastInsertId()
	if err != nil {
		return structs.User{}, err
	}

	if pictureURI != "" {
		query := "UPDATE users SET picture_uri=? WHERE id=?"
		if _, err := db.c.Exec(query, pictureURI, userID); err != nil {
			return structs.User{}, err
		}
	}

	return structs.User{
		ID:         int(userID),
		Username:   username,
		PictureURI: pictureURI,
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

func (db *appdbimpl) UpdateUserPicture(userID int, pictureURI string) error {
	query := "UPDATE users SET picture_uri=? WHERE id=?"
	_, err := db.c.Exec(query, pictureURI, userID)
	return err
}

/*
 * GET OPERATIONS
 */

func (db *appdbimpl) GetUser(userID int) (structs.User, error) {
	var user structs.User

	query := "SELECT id, username, COALESCE(picture_uri, '') FROM users WHERE id=?"
	err := db.c.QueryRow(query, userID).Scan(&user.ID, &user.Username, &user.PictureURI)
	return user, err
}

func (db *appdbimpl) GetUserPictureURI(userID int) (string, error) {
	var pictureURI string

	query := "SELECT COALESCE(picture_uri, '') FROM users WHERE id=?"
	err := db.c.QueryRow(query, userID).Scan(&pictureURI)
	return pictureURI, err
}

func (db *appdbimpl) GetUserByUsername(username string) (structs.User, error) {
	var user structs.User

	query := "SELECT id, username, COALESCE(picture_uri, '') FROM users WHERE username=?"
	err := db.c.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.PictureURI)
	return user, err
}

func (db *appdbimpl) GetUsers(username string) ([]structs.User, error) {
	query := "SELECT id, username, COALESCE(picture_uri, '') FROM users WHERE username LIKE CONCAT(?, '%')"
	rows, err := db.c.Query(query, username)
	if err != nil {
		return nil, err
	}

	var users []structs.User
	for rows.Next() {
		var user structs.User
		if err := rows.Scan(&user.ID, &user.Username, &user.PictureURI); err != nil {
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
