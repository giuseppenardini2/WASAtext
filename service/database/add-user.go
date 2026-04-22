package database

// Add a user and return its ID
func (db *appdbimpl) AddUser(user_id int, username string, picture string) (int, error) {
	var id int
	err := db.c.QueryRow("INSERT INTO User (username, picture) VALUES (?, ?) RETURNING id", username, picture).Scan(&id)
	return id, err
}
