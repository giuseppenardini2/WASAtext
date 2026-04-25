package database

// Given a user ID and a group ID, add the user to the group
func (db *appdbimpl) AddMemberToGroup(user_id int, group_id int) (error) {
	_, err := db.c.Exec("INSERT INTO user_group (user, group) VALUES (?, ?)", user_id, group_id)
	return err
}