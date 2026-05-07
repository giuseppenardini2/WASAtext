package database

import "wasatext/service/api/structs"

// Given a user ID and recipient user ID, return the user's messages with that user
func (db *appdbimpl) GetMessagesByRecipientUser(user_id int, recipient_user_id int) ([]structs.Message, error) {
	query := `SELECT id, body, sent_date, picture, sender_id, COALESCE(recipient_user_id, -1), COALESCE(recipient_group_id, -1), COALESCE(reply_message_id, -1), COALESCE(forwarded_message_id, -1)
		FROM messages
		WHERE (sender_id=? AND recipient_user_id=?) OR (sender_id=? AND recipient_user_id=?)
		ORDER BY sent_date DESC`
	rows, err := db.c.Query(query, user_id, recipient_user_id, recipient_user_id, user_id)
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
