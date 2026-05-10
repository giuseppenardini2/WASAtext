package database

// Add a reaction to a message
func (db *appdbimpl) AddReaction(userID int, messageID int, reaction string) error {
	query := "INSERT INTO users_react_messages (user_id, message_id, reaction) VALUES (?, ?)"
	_, err := db.c.Exec(query, userID, messageID, reaction)
	return err
}
