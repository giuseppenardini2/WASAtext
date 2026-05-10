package database

// Add a reply message to a user and return its ID
func (db *appdbimpl) AddReplyMessageToUser(body string, sentDate string, picture string, senderID int, recipientUserID int, replyMessageID int) (int64, error) {
	query := "INSERT INTO messages (body, sent_date, picture, sender_id, recipient_user_id, reply_message_id) VALUES (?, ?, ?, ?, ?, ?)"
	result, err := db.c.Exec(query, body, sentDate, picture, senderID, recipientUserID, replyMessageID)
	if err != nil {
		return -1, err
	}
	return result.LastInsertId()
}
