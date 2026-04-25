package database

// Given a group ID and a profile picture, update the profile picture of the group
func (db *appdbimpl) UpdateGroupPicture(id int, picture string) (error) {
	_, err := db.c.Exec("UPDATE Group SET picture=? WHERE id=?", picture, id)
	return err
}
