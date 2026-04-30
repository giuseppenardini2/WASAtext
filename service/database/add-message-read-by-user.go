package database

// Add a message read by a user
func (db *appdbimpl) AddMessageReadByUser(user_id int, message_id int) error {
	var query = "INSERT INTO user_reads_message () VALUES (user, message)"
	_, err := db.c.Exec(query, user_id, message_id)
	return err
}
