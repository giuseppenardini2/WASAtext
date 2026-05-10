package database

// Given a group ID, delete the group
func (db *appdbimpl) DeleteGroup(groupID int) error {
	query := "DELETE FROM groups WHERE id=?"
	_, err := db.c.Exec(query, groupID)
	return err
}
