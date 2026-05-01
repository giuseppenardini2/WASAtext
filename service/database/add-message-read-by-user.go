package database

// Add a message read by a user
func (db *appdbimpl) AddMessageReadByUser(user_id int, message_id int) error {
	var query = "INSERT INTO users_read_messages () VALUES (user_id, message_id)"
	_, err := db.c.Exec(query, user_id, message_id)
	return err
}
