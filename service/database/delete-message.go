package database

// Given a message ID, delete the message
func (db *appdbimpl) DeleteMessage(id int) error {
	var query = "DELETE FROM messages WHERE id=?"
	_, err := db.c.Exec(query, id)
	return err
}
