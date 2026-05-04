package database

// Add a group and return its ID
func (db *appdbimpl) AddGroup(name string, picture string, user_creator_id int) (int64, error) {
	query := "INSERT INTO groups (name, picture, user_creator_id) VALUES (?, ?, ?)"
	result, err := db.c.Exec(query, name, picture, user_creator_id)
	if err != nil {
		return -1, err
	}
	return result.LastInsertId()
}
