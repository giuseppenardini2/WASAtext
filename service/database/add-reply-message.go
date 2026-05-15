package database

import (
	"errors"
	"wasatext/service/structs"
)

// Add a reply message to a user and return its ID
func (db *appdbimpl) AddReplyMessage(body string, sentDate string, picture structs.Picture, senderID int, recipientID int, recipientType structs.RecipientType, replyMessageID int) (int64, error) {
	// Add the message based on the recipient type
	var query string
	switch recipientType {
	case structs.RecipientUser:
		query = "INSERT INTO messages (body, sent_date, picture, sender_id, recipient_user_id, reply_message_id) VALUES (?, ?, ?, ?, ?)"
	case structs.RecipientGroup:
		query = "INSERT INTO messages (body, sent_date, picture, sender_id, recipient_group_id, reply_message_id) VALUES (?, ?, ?, ?, ?)"
	default:
		return -1, errors.New("Recipient type error")
	}

	result, err := db.c.Exec(query, body, sentDate, senderID, recipientID, replyMessageID)
	if err != nil {
		return -1, err
	}

	// Add the picture of the message if it is present
	if !picture.IsNil {
		lastInsertID, err := result.LastInsertId()
		if err != nil {
			return -1, err
		}

		query := "UPDATE messages SET picture=? WHERE id=?"
		_, err = db.c.Exec(query, picture, lastInsertID)
		if err != nil {
			return -1, err
		}
	}

	return result.LastInsertId()
}
