package database

// Add a message to a group and return its ID
func (db *appdbimpl) AddReplyMessageToGroup(body string, sentDate string, picture string, senderID int, recipientGroupID int, replyMessageID int) (int64, error) {
	query := "INSERT INTO messages (body, sent_date, picture, sender_id, recipient_group_id, reply_message_id) VALUES (?, ?, ?, ?, ?, ?)"
	result, err := db.c.Exec(query, body, sentDate, picture, senderID, recipientGroupID, replyMessageID)
	if err != nil {
		return -1, err
	}
	return result.LastInsertId()
}
