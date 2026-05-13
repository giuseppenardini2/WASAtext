package database

// Given a user ID, return if the user is already used
func (db *appdbimpl) CheckUser(userID int) (bool, error) {
	var count int

	query := "SELECT COUNT(id) FROM users WHERE id=?"
	err := db.c.QueryRow(query, userID).Scan(&count)
	if err != nil {
		return false, err
	}
	return count != 0, err
}
