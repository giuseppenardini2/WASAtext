package database

// Given a user ID and a group ID, add the user to the group
func (db *appdbimpl) AddMemberToGroup(user_id int, group_id int) error {
	var query = "INSERT INTO users_groups (user_id, group_id) VALUES (?, ?)"
	_, err := db.c.Exec(query, user_id, group_id)
	return err
}
