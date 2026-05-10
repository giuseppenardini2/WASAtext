package database

// Add a group and return its ID
func (db *appdbimpl) AddGroup(name string, picture string, userCreatorID int) (int64, error) {
	query := "INSERT INTO groups (name, picture, user_creator_id) VALUES (?, ?, ?)"
	result, err := db.c.Exec(query, name, picture, userCreatorID)
	if err != nil {
		return -1, err
	}
	return result.LastInsertId()
}
