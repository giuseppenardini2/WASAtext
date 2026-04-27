package database

// Given a user ID and a profile picture, update the profile picture of the user
func (db *appdbimpl) UpdateUserPicture(id int, picture string) error {
	_, err := db.c.Exec("UPDATE User SET picture=? WHERE id=?", picture, id)
	return err
}
