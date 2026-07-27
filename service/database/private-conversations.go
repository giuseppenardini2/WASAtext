package database

import (
	"wasatext/service/structs"
)

/*
 * GET OPERATIONS
 */

func (db *appdbimpl) GetPrivateConversationMembers(privateConversationID int) ([]structs.User, error) {
	var users []structs.User

	query := `SELECT DISTINCT u.id, u.username, COALESCE(u.picture_uri, '')
	FROM private_conversations AS pc
	INNER JOIN users AS u ON pc.first_user_id=u.id OR pc.second_user_id=u.id
	WHERE pc.id=?`
	rows, err := db.c.Query(query, privateConversationID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var user structs.User
		if err := rows.Scan(&user.ID, &user.Username, &user.PictureURI); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

/*
 * ADD OPERATIONS
 */

func (db *appdbimpl) AddPrivateConversation(firstUserID int, secondUserID int) (int64, error) {
	tx, err := db.c.Begin()
	if err != nil {
		return -1, err
	}
	defer func() {
		if err := tx.Rollback(); err != nil {
			return
		}
	}()

	query1 := "INSERT INTO conversations DEFAULT VALUES"
	result, err := tx.Exec(query1)
	if err != nil {
		return -1, err
	}

	privateConversationID, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}

	if firstUserID > secondUserID {
		firstUserID, secondUserID = secondUserID, firstUserID
	}

	query2 := "INSERT INTO private_conversations (id, first_user_id, second_user_id) VALUES (?, ?, ?)"
	if _, err := tx.Exec(query2, privateConversationID, firstUserID, secondUserID); err != nil {
		return -1, err
	}

	if err := tx.Commit(); err != nil {
		return -1, err
	}

	return privateConversationID, nil
}

/*
 * CHECK OPERATIONS
 */

func (db *appdbimpl) CheckUserPrivateConversation(firstUserID int, secondUserID int) error {
	var count int

	query := "SELECT COUNT(id) FROM private_conversations WHERE (first_user_id=? AND second_user_id=?) OR (first_user_id=? AND second_user_id=?)"
	err := db.c.QueryRow(query, firstUserID, secondUserID, secondUserID, firstUserID).Scan(&count)
	if err != nil {
		return err
	}
	if count == 0 {
		return ErrNonexistentResource
	}
	return nil
}
