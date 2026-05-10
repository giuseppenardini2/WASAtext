package database

// Add a message to a group and return its ID
func (db *appdbimpl) AddMessageToGroup(body string, sentDate string, picture string, senderID int, recipientGroupID int) (int64, error) {
	query := "INSERT INTO messages (body, sent_date, picture, sender_id, recipient_group_id) VALUES (?, ?, ?, ?, ?)"
	result, err := db.c.Exec(query, body, sentDate, picture, senderID, recipientGroupID)
	if err != nil {
		return -1, err
	} //TODO: Aggiungere il messaggio solo se l'utente fa parte del gruppo
	return result.LastInsertId()
}
