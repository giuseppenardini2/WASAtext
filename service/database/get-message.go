package database

import "WASAtext/service/api/structs"

// Given a message ID, return the message
func (db *appdbimpl) GetMessage(id int) (structs.Message, error) {
	var message structs.Message
	err := db.c.QueryRow("SELECT id, body, date, picture, sender, recipient_user, recipient_group, reply_message, forwarded_message FROM Message WHERE id=?", id).Scan(
		&message.id, &message.date, &message.picture, &message.sender, &message.recipient_user, &message.recipient_group, &message.reply_message, &message.forwarded_message)
	if err != nil {
		return nil, err
	}
	return message, err
}