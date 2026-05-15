package database

import (
	"errors"
	"wasatext/service/structs"
)

// Add a message to a user and return its ID
func (db *appdbimpl) AddMessage(body string, sentDate string, picture string, senderID int, recipientID int, recipientType structs.RecipientType) (int64, error) {
	// Add the message based on the recipient type
	var query string
	switch recipientType {
	case structs.RecipientUser:
		query = "INSERT INTO messages (body, sent_date, sender_id, recipient_user_id) VALUES (?, ?, ?, ?)"
	case structs.RecipientGroup:
		query = "INSERT INTO messages (body, sent_date, sender_id, recipient_group_id) VALUES (?, ?, ?, ?)"
	default:
		return -1, errors.New("Recipient type error")
	}

	result, err := db.c.Exec(query, body, sentDate, picture, senderID, recipientID)
	if err != nil {
		return -1, err
	}

	// Add the picture of the message if it is present
	if picture != "" {
		lastInsertId, err := result.LastInsertId()
		if err != nil {
			return -1, err
		}

		query := "UPDATE messages SET picture=? WHERE id=?"
		_, err = db.c.Exec(query, picture, lastInsertId)
		if err != nil {
			return -1, err
		}
	}

	return result.LastInsertId()
}
