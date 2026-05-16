package database

import "wasatext/service/structs"

// Add a forwarded message to a user and return its ID
func (db *appdbimpl) AddForwardedMessage(originalMessageID int, senderID int, recipientID int, recipientType structs.RecipientType, sentDate string) (int64, error) {
	// Get the body and picture of the message to forward to
	var body, picture string

	query1 := `SELECT body, COALESCE(picture, "") FROM messages WHERE id=?`
	err := db.c.QueryRow(query1, originalMessageID).Scan(&body, &picture)
	if err != nil {
		return -1, err
	}

	// Create a new message with the same body and picture of the message to forward to
	var query2 string
	switch recipientType {
	case structs.RecipientUser:
		query2 = "INSERT INTO messages (body, sent_date, sender_id, recipient_user_id, forwarded_message_id) VALUES (?, ?, ?, ?, ?)"
	case structs.RecipientGroup:
		query2 = "INSERT INTO messages (body, sent_date, sender_id, recipient_group_id, forwarded_message_id) VALUES (?, ?, ?, ?, ?)"
	default:
		return -1, nil
	}

	result, err := db.c.Exec(query2, body, sentDate, senderID, recipientID, originalMessageID)
	if err != nil {
		return -1, err
	}

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
