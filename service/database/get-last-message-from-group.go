package database

import "wasatext/service/api/structs"

// Given a user ID and recipient group ID, return the last message with that group
func (db *appdbimpl) GetLastMessageFromGroup(userID int, recipientGroupID int) (structs.Message, error) {
	var message structs.Message

	query := `SELECT m.id, m.body, m.sent_date, m.picture, m.sender_id, COALESCE(m.recipient_user_id, -1), COALESCE(m.recipient_group_id, -1), COALESCE(m.reply_message_id, -1), COALESCE(m.forwarded_message_id, -1)
		FROM messages AS m, users_groups AS ug
		WHERE ug.user_id=? AND ug.group_id=? AND m.recipient_group_id=?
		ORDER BY sent_date DESC LIMIT 1`
	err := db.c.QueryRow(query, userID, recipientGroupID).Scan(&message.ID, &message.SentDate, &message.Picture, &message.SenderID, &message.RecipientUserID, &message.RecipientGroupID, &message.ReplyMessageID, &message.ForwardedMessageID)
	return message, err
}
