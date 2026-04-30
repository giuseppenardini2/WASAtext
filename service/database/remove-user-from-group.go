package database

// Given a user ID and group ID, remove the user from that group
func (db *appdbimpl) RemoveUserFromGroup(user_id int, group_id int) error {
	var query = "DELETE FROM user_group WHERE user=? AND group=?"
	_, err := db.c.Exec(query, user_id, group_id)
	return err
}
