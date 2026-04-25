package database

// Add a message and return its ID
//TODO: Vedere se modificare il booleano per inviare un messaggio al gruppo
func (db *appdbimpl) AddMessage(body string, date string, picture string, sender int, recipient int, is_recipient_user bool, reply_message int, forwarded_message int) (int64, error) {
	if is_recipient_user==true {
		result, err := db.c.Exec("INSERT INTO Message (body, date, picture, sender, recipient_user, reply_message, forwarded_message) VALUES (?, ?, ?, ?, ?, ?, ?)", body, date, picture, sender, recipient, forwarded_message)
		if err != nil {
			return -1, err
		}
		return result.LastInsertId(), err
	}
	result, err := db.c.Exec("INSERT INTO Message (body, date, picture, sender, recipient_group, reply_message, forwarded_message) VALUES (?, ?, ?, ?, ?, ?, ?)", body, date, picture, sender, recipient, forwarded_message)
	if err != nil {
		return -1, err
	}
	return result.LastInsertId(), err
}