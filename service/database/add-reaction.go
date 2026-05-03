package database

// Add a reaction to a message
func (db *appdbimpl) AddReaction(user_id int, message_id int, reaction string) error {
	query := "INSERT INTO users_react_messages (user_id, message_id, reaction) VALUES (?, ?)"
	_, err := db.c.Exec(query, user_id, message_id, reaction)
	return err
}
