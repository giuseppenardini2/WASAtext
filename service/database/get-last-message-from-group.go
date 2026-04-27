package database

import "WASAtext/service/api/structs"

// Given a user ID and recipient group ID, return the last message with that group
func (db *appdbimpl) GetLastMessageFromUser(user_id int, recipient_group int) (structs.Message, error) {
	var message structs.Message
	err := db.c.QueryRow("SELECT id, body, date, picture, sender, recipient_user, recipient_group, reply_message, forwarded_message FROM Message 
		WHERE sender=? AND recipient_group=?
		ORDER BY date DESC LIMIT 1", user_id, recipient_group).Scan(
		&message.id, &message.date, &message.picture, &message.sender, &message.recipient_user, &message.recipient_group, &message.reply_message, &message.forwarded_message)
	if err != nil {
		return nil, err
	}
	return message, err
}