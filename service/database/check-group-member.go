package database

// Given a user ID and a group ID, return if the user is a member of the group
func (db *appdbimpl) CheckGroupMember(userID int, groupID int) (bool, error) {
	var count int

	query := "SELECT COUNT(user_id) FROM users_groups WHERE user_id=? AND group_id=?"
	err := db.c.QueryRow(query, userID, groupID).Scan(&count)
	if err != nil {
		return false, err
	}
	return count != 0, err
}
