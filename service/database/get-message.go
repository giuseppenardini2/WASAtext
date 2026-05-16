package database

import "wasatext/service/structs"

// Given a message ID, return the message
func (db *appdbimpl) GetMessage(messageID int) (structs.Message, error) {
	var message structs.Message

	query1 := `SELECT id, body, sent_date, picture, sender_id, COALESCE(recipient_user_id, -1), COALESCE(recipient_group_id, -1), COALESCE(reply_message_id, -1), COALESCE(forwarded_message_id, -1)
		FROM messages
		WHERE id=?`
	err := db.c.QueryRow(query1, messageID).Scan(&message.ID, &message.SentDate, &message.Picture, &message.SenderID, &message.RecipientUserID, &message.RecipientGroupID, &message.ReplyMessageID, &message.ForwardedMessageID)
	return message, err
}
