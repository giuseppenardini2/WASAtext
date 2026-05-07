package database

// Add a message to a group and return its ID
func (db *appdbimpl) AddMessageToGroup(body string, sent_date string, picture string, sender_id int, recipient_group_id int) (int64, error) {
	query := "INSERT INTO messages (body, sent_date, picture, sender_id, recipient_group_id) VALUES (?, ?, ?, ?, ?)"
	result, err := db.c.Exec(query, body, sent_date, picture, sender_id, recipient_group_id)
	if err != nil {
		return -1, err
	} //TODO: Aggiungere il messaggio solo se l'utente fa parte del gruppo
	return result.LastInsertId()
}
