package database

// Given a user ID and a profile picture, update the profile picture of the user
func (db *appdbimpl) UpdateUserPicture(id int, picture string) error {
	var query = "UPDATE User SET picture=? WHERE id=?"
	_, err := db.c.Exec(query, picture, id)
	return err
}
