package database

// Add a forwarded message to a user and return its ID
func (db *appdbimpl) AddForwardedMessageToUser(originalMessageID int, senderID int, recipientUserID int, sentDate string) (int64, error) {
	//Get the body and picture of the message to forward to
	var body, picture string

	query1 := "SELECT body, picture FROM messages WHERE id=?"
	err := db.c.QueryRow(query1, originalMessageID).Scan(&body, &picture)
	if err != nil {
		return -1, err
	}

	//Create a new message with the same body and picture of the message to forward to
	query2 := "INSERT INTO messages (body, sent_date, picture, sender_id, recipient_user_id, forwarded_message_id) VALUES (?, ?, ?, ?, ?, ?)"
	result, err := db.c.Exec(query2, body, sentDate, picture, senderID, recipientUserID, originalMessageID)
	if err != nil {
		return -1, err
	}
	return result.LastInsertId()
}
