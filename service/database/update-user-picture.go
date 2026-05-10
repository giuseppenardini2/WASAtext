package database

// Given a user ID and a profile picture, update the profile picture of the user
func (db *appdbimpl) UpdateUserPicture(userID int, picture string) error {
	query := "UPDATE users SET picture=? WHERE id=?"
	_, err := db.c.Exec(query, picture, userID)
	return err
}
