package database

import (
	"errors"
	"wasatext/service/structs"
)

/*
 * ADD OPERATIONS
 */

func (db *appdbimpl) AddReadCheckmark(userID int, messageID int, conversationID int) error {
	query := "INSERT INTO users_read_messages (user_id, message_id, conversation_id) VALUES (?, ?, ?)"
	_, err := db.c.Exec(query, userID, messageID, conversationID)
	return err
}

/*
 * CHECK OPERATIONS
 */

func (db *appdbimpl) CheckReadMessage(messageID int, conversationID int) error {
	// Check if the recipient of the message is a user or a group
	conversationType, err := db.GetConversationType(conversationID)
	if err != nil {
		return err
	}

	// If the recipient is a user, check if that user read the message. Else if it is a group, check if all group members read the message
	var usersReadMessageNumber int
	query2 := `SELECT COUNT(message_id)
			FROM read_checkmarks
			WHERE message_id=? AND conversation_id=?`
	err = db.c.QueryRow(query2, messageID, conversationID).Scan(&usersReadMessageNumber)
	if err != nil || usersReadMessageNumber == 0 {
		return err
	}
	switch conversationType {
	case structs.PrivateConversationType:
		if usersReadMessageNumber == 0 {
			return ErrNonexistentResource
		}
		return nil
	case structs.GroupConversationType:
		var groupMembersNumber int
		query3 := `SELECT COUNT(gcm.user_id)
			FROM group_conversations_members AS gcm
			INNER JOIN messages AS m ON gcm.group_conversation_id=m.conversation_id
			WHERE m.id=? AND m.conversation_id=?`
		err = db.c.QueryRow(query3, messageID, conversationID).Scan(&groupMembersNumber)
		if err != nil {
			return err
		}
		if usersReadMessageNumber != groupMembersNumber {
			return ErrNonexistentResource
		}
		return nil
	default:
		return errors.New("An error occurs in the recipient type")
	}
}
