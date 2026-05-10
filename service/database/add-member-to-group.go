package database

// Given a user ID and a group ID, add the user to the group
func (db *appdbimpl) AddMemberToGroup(userID int, groupID int) error {
	query := "INSERT INTO users_groups (user_id, group_id) VALUES (?, ?)"
	_, err := db.c.Exec(query, userID, groupID)
	return err
}
