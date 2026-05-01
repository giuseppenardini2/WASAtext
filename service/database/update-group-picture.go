package database

// Given a group ID and a profile picture, update the profile picture of the group
func (db *appdbimpl) UpdateGroupPicture(id int, picture string) error {
	var query = "UPDATE groups SET picture=? WHERE id=?"
	_, err := db.c.Exec(query, picture, id)
	return err
}
