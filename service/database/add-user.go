package database

// Add a user and return its ID
func (db *appdbimpl) AddUser(username string, picture string) (int64, error) {
	query := "INSERT INTO users (username, picture) VALUES (?, ?)"
	result, err := db.c.Exec(query, username, picture)
	if err != nil {
		return -1, err
	}
	return result.LastInsertId()
}
