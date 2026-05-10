package database

// Given a user ID and message ID, delete the user's reaction to that message
func (db *appdbimpl) DeleteReaction(userID int, messageID int) error {
	query := "DELETE FROM users_react_messages WHERE user_id=? AND message_id=?"
	_, err := db.c.Exec(query, userID, messageID)
	return err
}
