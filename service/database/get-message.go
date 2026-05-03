package database

import "wasatext/service/api/structs"

// Given a message ID, return the message
func (db *appdbimpl) GetMessage(id int) (structs.Message, error) {
	var message structs.Message

	query := "SELECT id, body, date, picture, sender_id, recipient_user_id, recipient_group_id, reply_message_id, forwarded_message_id FROM messages WHERE id=?"
	err := db.c.QueryRow(query, id).Scan(&message.ID, &message.Date, &message.Picture, &message.SenderID, &message.RecipientUserID, &message.RecipientGroupID, &message.ReplyMessageID, &message.ForwardedMessageID)
	return message, err
}
