package database

// Add a message to a group and return its ID
func (db *appdbimpl) AddReplyMessageToGroup(body string, date string, picture string, sender_id int, recipient_group_id int, reply_message_id int) (int64, error) {
	query := "INSERT INTO messages (body, date, picture, sender_id, recipient_group_id, reply_message_id) VALUES (?, ?, ?, ?, ?, ?)"
	result, err := db.c.Exec(query, body, date, picture, sender_id, recipient_group_id, reply_message_id)
	if err != nil {
		return -1, err
	}
	return result.LastInsertId()
}
