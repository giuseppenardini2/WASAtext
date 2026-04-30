package database

// Add a message to a group and return its ID
func (db *appdbimpl) AddMessageToGroup(body string, date string, picture string, sender int, recipient_group int, reply_message int) (int64, error) {
	var query = "INSERT INTO Message (body, date, picture, sender, recipient_group, reply_message) VALUES (?, ?, ?, ?, ?, ?)"
	result, err := db.c.Exec(query, body, date, picture, sender, recipient_group, reply_message)
	if err != nil {
		return -1, err
	}
	return result.LastInsertId()
}
