package database

// Given a user ID and a group, return if the user is a member of the group
func (db *appdbimpl) CheckGroupMember(user_id int, group_id int) (bool, error) {
	var count int

	query := "SELECT COUNT(user_id) FROM users_groups WHERE user_id=? AND group_id=?"
	err := db.c.QueryRow(query, user_id, group_id).Scan(&count)
	if err != nil {
		return false, err
	}
	return count != 0, err
}
