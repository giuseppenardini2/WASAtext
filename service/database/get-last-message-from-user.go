package database

import "wasatext/service/api/structs"

// Given a user ID and recipient user ID, return the last message with that user
func (db *appdbimpl) GetLastMessageFromUser(userID int, recipientUserID int) (structs.Message, error) {
	var message structs.Message

	query := `SELECT id, body, sent_date, picture, sender_id, COALESCE(recipient_user_id, -1), COALESCE(recipient_group_id, -1), COALESCE(reply_message_id, -1), COALESCE(forwarded_message_id, -1)
		FROM messages
		WHERE (sender_id=? AND recipient_user_id=?) OR (sender_id=? AND recipient_user_id=?)
		ORDER BY sent_date DESC LIMIT 1`
	err := db.c.QueryRow(query, userID, recipientUserID).Scan(&message.ID, &message.SentDate, &message.Picture, &message.SenderID, &message.RecipientUserID, &message.RecipientGroupID, &message.ReplyMessageID, &message.ForwardedMessageID)
	return message, err
}
