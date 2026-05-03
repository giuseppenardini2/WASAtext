package database

// Add a forwarded message to a user and return its ID
func (db *appdbimpl) AddForwardedMessageToUser(original_message_id int, sender_id int, recipient_user_id int, date string) (int64, error) {
	//Get the body and picture of the message to forward to
	var body, picture string

	query1 := "SELECT body, picture FROM messages WHERE id=?"
	err := db.c.QueryRow(query1, original_message_id).Scan(&body, &picture)
	if err != nil {
		return -1, err
	}

	//Create a new message with the same body and picture of the message to forward to
	query2 := "INSERT INTO messages (body, date, picture, sender_id, recipient_user_id, forwarded_message_id) VALUES (?, ?, ?, ?, ?, ?)"
	result, err := db.c.Exec(query2, body, date, picture, sender_id, recipient_user_id, original_message_id)
	if err != nil {
		return -1, err
	}
	return result.LastInsertId()
}
