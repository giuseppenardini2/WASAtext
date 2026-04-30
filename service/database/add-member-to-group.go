package database

// Given a user ID and a group ID, add the user to the group
func (db *appdbimpl) AddMemberToGroup(user_id int, group_id int) error {
	var query = "INSERT INTO user_group (user, group) VALUES (?, ?)"
	_, err := db.c.Exec(query, user_id, group_id)
	return err
}
