package database

import "wasatext/service/api/structs"

// Given a user ID and recipient group ID, return the last message with that group
func (db *appdbimpl) GetLastMessageFromGroup(user_id int, recipient_group_id int) (structs.Message, error) {
	var message structs.Message

	query := `SELECT id, body, date, picture, sender_id, COALESCE(recipient_user_id, -1), COALESCE(recipient_group_id, -1), COALESCE(reply_message_id, -1), COALESCE(forwarded_message_id, -1)
		FROM messages
		WHERE sender_id=? AND recipient_group_id=?
		ORDER BY date DESC LIMIT 1`
	err := db.c.QueryRow(query, user_id, recipient_group_id).Scan(&message.ID, &message.Date, &message.Picture, &message.SenderID, &message.RecipientUserID, &message.RecipientGroupID, &message.ReplyMessageID, &message.ForwardedMessageID)
	return message, err
}
