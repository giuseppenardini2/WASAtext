package database

// Add a message to a user and return its ID
func (db *appdbimpl) AddMessageToUser(body string, date string, picture string, sender_id int, recipient_user_id int) (int64, error) {
	query := "INSERT INTO messages (body, date, picture, sender_id, recipient_user_id) VALUES (?, ?, ?, ?, ?)"
	result, err := db.c.Exec(query, body, date, picture, sender_id, recipient_user_id)
	if err != nil {
		return -1, err
	}
	return result.LastInsertId()
}
