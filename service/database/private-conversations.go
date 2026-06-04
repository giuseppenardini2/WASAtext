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

/*
 * CHECK OPERATIONS
 */

func (db *appdbimpl) CheckUserPrivateConversation(firstUserID int, secondUserID int) error {
	var count int

	query := "SELECT COUNT(id) FROM private_conversations WHERE (first_user_id=? AND second_user_id=?) OR (first_user_id=? AND second_user_id=?) "
	err := db.c.QueryRow(query, firstUserID, secondUserID, secondUserID, firstUserID).Scan(&count)
	if err != nil {
		return err
	}
	if count == 0 {
		return ErrNonexistentResource
	}
	return nil
}
