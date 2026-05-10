package database

// Add a message read by a user
func (db *appdbimpl) AddMessageReadByUser(userID int, messageID int) error {
	query := "INSERT INTO users_read_messages (user_id, message_id) VALUES (?, ?)"
	_, err := db.c.Exec(query, userID, messageID)
	return err
}
