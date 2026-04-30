package database

// Add a forwarded message to a user and return its ID
func (db *appdbimpl) AddForwardedMessageToUser(original_message_id int, sender int, recipient_user int, date string) (int64, error) {
	//Get the body and picture of the message to forward to
	var body, picture string
	err := db.c.QueryRow("SELECT body, picture FROM Message WHERE id=?", original_message_id).Scan(&body, &picture)
	if err != nil {
		return -1, err
	}

	//Create a new message with the same body and picture of the message to forward to
	result, err := db.c.Exec("INSERT INTO Message (body, date, picture, sender, recipient_user, forwarded_message) VALUES (?, ?, ?, ?, ?, ?)", body, date, picture, sender, recipient_user, original_message_id)
	if err != nil {
		return -1, err
	}
	return result.LastInsertId()
}
