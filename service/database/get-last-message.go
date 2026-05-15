package database

import (
	"errors"
	"wasatext/service/structs"
)

// Given a user ID and recipient group ID, return the last message with that group
func (db *appdbimpl) GetLastMessage(userID int, recipientID int, recipientType structs.RecipientType) (structs.Message, error) {
	var message structs.Message
	var err error

	switch recipientType {
	case structs.RecipientUser:
		query := `SELECT id, body, sent_date, picture, sender_id, COALESCE(recipient_user_id, -1), COALESCE(recipient_group_id, -1), COALESCE(reply_message_id, -1), COALESCE(forwarded_message_id, -1)
		FROM messages
		WHERE (sender_id=? AND recipient_user_id=?) OR (sender_id=? AND recipient_user_id=?)
		ORDER BY sent_date DESC LIMIT 1`
		err = db.c.QueryRow(query, userID, recipientID).Scan(&message.ID, &message.SentDate, &message.Picture, &message.SenderID, &message.RecipientUserID, &message.RecipientGroupID, &message.ReplyMessageID, &message.ForwardedMessageID)
	case structs.RecipientGroup:
		query := `SELECT m.id, m.body, m.sent_date, m.picture, m.sender_id, COALESCE(m.recipient_user_id, -1), COALESCE(m.recipient_group_id, -1), COALESCE(m.reply_message_id, -1), COALESCE(m.forwarded_message_id, -1)
		FROM messages AS m, users_groups AS ug
		WHERE ug.user_id=? AND ug.group_id=? AND m.recipient_group_id=?
		ORDER BY sent_date DESC LIMIT 1`
		err = db.c.QueryRow(query, userID, recipientID).Scan(&message.ID, &message.SentDate, &message.Picture, &message.SenderID, &message.RecipientUserID, &message.RecipientGroupID, &message.ReplyMessageID, &message.ForwardedMessageID)
	default:
		return message, errors.New("Recipient type error")
	}
	return message, err
}
