package database

// Add a reply message to a user and return its ID
func (db *appdbimpl) AddReplyMessageToUser(body string, date string, picture string, sender_id int, recipient_user_id int, reply_message_id int) (int64, error) {
	query := "INSERT INTO messages (body, date, picture, sender_id, recipient_user_id, reply_message_id) VALUES (?, ?, ?, ?, ?, ?)"
	result, err := db.c.Exec(query, body, date, picture, sender_id, recipient_user_id, reply_message_id)
	if err != nil {
		return -1, err
	}
	return result.LastInsertId()
}
