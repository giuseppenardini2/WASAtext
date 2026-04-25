package database

// Add a user and return its ID
func (db *appdbimpl) AddUser(username string, picture string) (int64, error) {
	result, err := db.c.Exec("INSERT INTO User (username, picture) VALUES (?, ?)", username, picture)
	if err != nil {
		return -1, err
	}
	return result.LastInsertId(), err
}