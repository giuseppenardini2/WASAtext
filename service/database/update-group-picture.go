package database

// Given a group ID and a profile picture, update the profile picture of the group
func (db *appdbimpl) UpdateGroupPicture(groupID int, picture string) error {
	query := "UPDATE groups SET picture=? WHERE id=?"
	_, err := db.c.Exec(query, picture, groupID)
	return err
}
