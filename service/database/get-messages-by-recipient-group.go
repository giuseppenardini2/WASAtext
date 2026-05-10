package database

import "wasatext/service/structs"

// Given a user ID and recipient group ID, return the user's messages with that group
func (db *appdbimpl) GetMessagesByRecipientGroup(userID int, recipientGroupID int) ([]structs.Message, error) {
	query := `SELECT m.id, m.body, m.sent_date, m.picture, m.sender_id, COALESCE(m.recipient_user_id, -1), COALESCE(m.recipient_group_id, -1), COALESCE(m.reply_message_id, -1), COALESCE(m.forwarded_message_id, -1)
		FROM messages AS m, users_groups AS ug
		WHERE ug.user_id=? AND ug.group_id=? AND m.recipient_group_id=?
		ORDER BY sent_date DESC`
	rows, err := db.c.Query(query, userID, recipientGroupID, recipientGroupID)
	if err != nil {
		return nil, err
	}

	var messages []structs.Message
	for rows.Next() {
		var message structs.Message
		if err := rows.Scan(&message.ID, &message.Body, &message.SentDate, &message.Picture, &message.SenderID, &message.RecipientUserID, &message.RecipientGroupID, &message.ReplyMessageID, &message.ForwardedMessageID); err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return messages, err
}
