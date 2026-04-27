package database

// Add a reaction to a message
func (db *appdbimpl) AddReaction(user_id int, message_id int, reaction string) error {
	_, err := db.c.Exec("INSERT INTO user_reacts_message (user, message, reaction) VALUES (?, ?)", user_id, message_id, reaction)
	return err
}