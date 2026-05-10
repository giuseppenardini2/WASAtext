package database

// Add a message to a user and return its ID
func (db *appdbimpl) AddMessageToUser(body string, sentDate string, picture string, senderID int, recipientUserID int) (int64, error) {
	query := "INSERT INTO messages (body, sent_date, picture, sender_id, recipient_user_id) VALUES (?, ?, ?, ?, ?)"
	result, err := db.c.Exec(query, body, sentDate, picture, senderID, recipientUserID)
	if err != nil {
		return -1, err
	}
	return result.LastInsertId()
}
