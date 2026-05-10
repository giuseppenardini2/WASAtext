package database

// Given a user ID and group ID, remove the user from that group
func (db *appdbimpl) RemoveUserFromGroup(userID int, groupID int) error {
	query := "DELETE FROM users_groups WHERE user_id=? AND group_id=?"
	_, err := db.c.Exec(query, userID, groupID)
	return err
}
