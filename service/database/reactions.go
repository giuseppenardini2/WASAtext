package database

/*
 * ADD OPERATIONS
 */
// Add a reaction to a message
func (db *appdbimpl) AddReaction(userID int, messageID int, conversationID int, reaction string) error {
	query := "INSERT INTO reactions (user_id, message_id, conversation_id, reaction) VALUES (?, ?, ?, ?)"
	_, err := db.c.Exec(query, userID, messageID, conversationID, reaction)
	return err
}

/*
 * DELETE OPERATIONS
 */
// Given a user ID and message ID, delete the user's reaction to that message
func (db *appdbimpl) DeleteReaction(userID int, messageID int, conversationID int) error {
	query := "DELETE FROM reactions WHERE user_id=? AND message_id=? AND conversation_id=?"
	_, err := db.c.Exec(query, userID, messageID, conversationID)
	return err
}
