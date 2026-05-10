package database

// Given a username, return if the username is already used
func (db *appdbimpl) CheckUserUsername(username string) (bool, error) {
	var count int

	query := "SELECT COUNT(id) FROM users WHERE username=?"
	err := db.c.QueryRow(query, username).Scan(&count)
	if err != nil {
		return false, err
	}
	return count != 0, err
}
