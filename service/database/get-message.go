package database

import "wasatext/service/api/structs"

// Given a message ID, return the message
func (db *appdbimpl) GetMessage(id int) (structs.Message, error) {
	var message structs.Message

	var query = "SELECT id, body, date, picture, sender, recipient_user, recipient_group, reply_message, forwarded_message FROM messages WHERE id=?"
	err := db.c.QueryRow(query, id).Scan(&message.Id, &message.Date, &message.Picture, &message.Sender, &message.RecipientUser, &message.RecipientGroup, &message.ReplyMessage, &message.ForwardedMessage)
	return message, err
}
