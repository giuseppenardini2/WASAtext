package database

import (
	"errors"
	"wasatext/service/structs"
)

/*
 * ADD OPERATIONS
 */

func (db *appdbimpl) AddReadCheckmark(userID int, messageID int, conversationID int) error {
	query := "INSERT INTO read_checkmarks (user_id, message_id, conversation_id) VALUES (?, ?, ?)"
	_, err := db.c.Exec(query, userID, messageID, conversationID)
	return err
}

func (db *appdbimpl) AddReadCheckmarks(userID int, conversationID int, messages []structs.Message) error {
	query1 := "SELECT message_id FROM read_checkmarks WHERE conversation_id=? AND user_id=?"

	rows, err := db.c.Query(query1, conversationID, userID, userID)
	if err != nil {
		return err
	}
	defer rows.Close()

	messagesReadByUser := make(map[int]struct{})
	for rows.Next() {
		var messageID int
		if err := rows.Scan(&messageID); err != nil {
			return err
		}
		messagesReadByUser[messageID] = struct{}{}
	}
	if err := rows.Err(); err != nil {
		return err
	}

	for i := range messages {
		if messages[i].SenderID != userID {
			_, isMessageRead := messagesReadByUser[messages[i].ID]
			if !isMessageRead {
				if err := db.AddReadCheckmark(userID, messages[i].ID, conversationID); err != nil {
					return err
				}
			}
		}
		err := db.CheckReadMessage(messages[i].ID, conversationID)
		switch {
		case errors.Is(err, ErrNonexistentResource):
			messages[i].Status = structs.SentMessageStatus
		case err == nil:
			messages[i].Status = structs.ReadMessageStatus
		default:
			return err
		}
	}

	return nil
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
	query2 := `SELECT COUNT(user_id)
			FROM read_checkmarks
			WHERE message_id=? AND conversation_id=?`
	if err := db.c.QueryRow(query2, messageID, conversationID).Scan(&usersReadMessageNumber); err != nil {
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
		if err := db.c.QueryRow(query3, messageID, conversationID).Scan(&groupMembersNumber); err != nil {
			return err
		}
		if usersReadMessageNumber != groupMembersNumber-1 {
			return ErrNonexistentResource
		}
		return nil
	default:
		return errors.New("An error occurs in the recipient type")
	}
}
