package database

// Add a message read by a user
func (db *appdbimpl) AddMessageReadByUser(user_id int, message_id int) error {
	_, err := db.c.Exec("INSERT INTO user_reads_message () VALUES (user, message)", user_id, message_id)
	return err
}