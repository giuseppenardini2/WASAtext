package database

// Given a message ID, delete the message
func (db *appdbimpl) DeleteMessage(messageID int) error {
	query := "DELETE FROM messages WHERE id=?"
	_, err := db.c.Exec(query, messageID)
	return err
}
