package database

// Given a user ID and a username, update the username of the user
func (db *appdbimpl) UpdateUserUsername(id int, username string) error {
	query := "UPDATE users SET username=? WHERE id=?"
	_, err := db.c.Exec(query, username, id)
	return err
}
