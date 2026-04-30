package database

// Given a group ID, delete the group
func (db *appdbimpl) DeleteGroup(id int) error {
	var query = "DELETE FROM Group WHERE id=?"
	_, err := db.c.Exec(query, id)
	return err
}
