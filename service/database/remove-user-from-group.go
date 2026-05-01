package database

// Given a user ID and group ID, remove the user from that group
func (db *appdbimpl) RemoveUserFromGroup(user_id int, group_id int) error {
	var query = "DELETE FROM users_groups WHERE user_id=? AND group_id=?"
	_, err := db.c.Exec(query, user_id, group_id)
	return err
}
