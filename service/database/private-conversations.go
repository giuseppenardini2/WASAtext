package database

/*
 * ADD OPERATIONS
 */
func (db *appdbimpl) AddPrivateConversation(firstUserID int, secondUserID int) (int64, error) {
	privateConversationID, err := db.AddConversation()
	if err != nil {
		return -1, err
	}

	query := "INSERT INTO private_conversations (id, first_user_id, second_user_id) VALUES (?, ?, ?)"
	_, err = db.c.Exec(query, privateConversationID, firstUserID, secondUserID)
	if err != nil {
		return -1, err
	}
	return privateConversationID, nil
}
