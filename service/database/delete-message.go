package database

// Given a message ID, delete the message
func (db *appdbimpl) DeleteMessage(id int) error {
	_, err := db.c.Exec("DELETE FROM Message WHERE id=?", id)
	return err
}