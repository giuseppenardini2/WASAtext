package database

// Add a group and return its ID
func (db *appdbimpl) AddGroup(name string, picture string, user_creator int) (int64, error) {
	var query = "INSERT INTO groups (name, picture, user_creator) VALUES (?, ?)"
	result, err := db.c.Exec(query, name, picture, user_creator)
	if err != nil {
		return -1, err
	}
	return result.LastInsertId()
}
