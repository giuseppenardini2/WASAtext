package database

// Add a message to a user and return its ID
func (db *appdbimpl) AddMessageToUser(body string, date string, picture string, sender int, recipient_user int, reply_message int) (int64, error) {
	var query = "INSERT INTO messages (body, date, picture, sender, recipient_user, reply_message) VALUES (?, ?, ?, ?, ?, ?)"
	result, err := db.c.Exec(query, body, date, picture, sender, recipient_user, reply_message)
	if err != nil {
		return -1, err
	}
	return result.LastInsertId()
}
