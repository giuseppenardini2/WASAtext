package database

import "wasatext/service/api/structs"

// Given a user ID and recipient group ID, return the last message with that group
func (db *appdbimpl) GetLastMessageFromGroup(user_id int, recipient_group int) (structs.Message, error) {
	var message structs.Message

	var query = `SELECT id, body, date, picture, sender, recipient_user, recipient_group, reply_message, forwarded_message FROM Message 
		WHERE sender=? AND recipient_group=?
		ORDER BY date DESC LIMIT 1`
	err := db.c.QueryRow(query, user_id, recipient_group).Scan(&message.Id, &message.Date, &message.Picture, &message.Sender, &message.RecipientUser, &message.RecipientGroup, &message.ReplyMessage, &message.ForwardedMessage)
	return message, err
}
