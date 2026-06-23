package database

import (
	"wasatext/service/structs"
)

/*
 * ADD OPERATIONS
 */

func (db *appdbimpl) AddGroupConversationMember(userID int, groupConversationID int) error {
	query := "INSERT INTO users_group_conversations (user_id, group_conversation_id) VALUES (?, ?)"
	_, err := db.c.Exec(query, userID, groupConversationID)
	return err
}

/*
 * DELETE OPERATIONS
 */

func (db *appdbimpl) DeleteGroupConversationMember(userID int, groupConversationID int) error {
	query := "DELETE FROM users_group_conversations WHERE user_id=? AND group_conversation_id=?"
	_, err := db.c.Exec(query, userID, groupConversationID)
	return err
}

/*
 * GET OPERATIONS
 */

func (db *appdbimpl) GetGroupConversartionMembers(groupConversationsID int) ([]structs.User, error) {
	query := `SELECT u.id, u.username, u.picture_uri 
		FROM users AS u, group_conversations_members AS gcm 
		WHERE u.id=gcm.user_id AND gcm.group_conversation_id=?`
	rows, err := db.c.Query(query, groupConversationsID)
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

func (db *appdbimpl) CheckGroupConversationMember(userID int, groupConversationID int) error {
	var count int

	query := "SELECT COUNT(id) FROM group_conversations_members WHERE user_id=? AND group_conversation_id=?"
	err := db.c.QueryRow(query, userID, groupConversationID).Scan(&count)
	if err != nil {
		return err
	}
	if count == 0 {
		return ErrNonexistentResource
	}
	return nil
}
