package database

import (
	"database/sql"
	"errors"
	"wasatext/service/structs"
)

// Given a user ID and recipient user ID, return the user's messages with that user
func (db *appdbimpl) GetMessagesByRecipient(userID int, recipientID int, recipientType structs.RecipientType) ([]structs.Message, error) {
	var rows *sql.Rows

	switch recipientType {
	case structs.RecipientUser:
		query := `SELECT id, body, sent_date, picture, sender_id, COALESCE(recipient_user_id, -1), COALESCE(recipient_group_id, -1), COALESCE(reply_message_id, -1), COALESCE(forwarded_message_id, -1)
		FROM messages
		WHERE (sender_id=? AND recipient_user_id=?) OR (sender_id=? AND recipient_user_id=?)
		ORDER BY sent_date DESC`
		var err error
		rows, err = db.c.Query(query, userID, recipientID, recipientID, userID)
		if err != nil {
			return nil, err
		}
	case structs.RecipientGroup:
		query := `SELECT m.id, m.body, m.sent_date, m.picture, m.sender_id, COALESCE(m.recipient_user_id, -1), COALESCE(m.recipient_group_id, -1), COALESCE(m.reply_message_id, -1), COALESCE(m.forwarded_message_id, -1)
		FROM messages AS m, users_groups AS ug
		WHERE ug.user_id=? AND ug.group_id=? AND m.recipient_group_id=?
		ORDER BY sent_date DESC`
		var err error
		rows, err = db.c.Query(query, userID, recipientID, recipientID)
		if err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("Recipient type error")
	}

	var messages []structs.Message
	for rows.Next() {
		var message structs.Message
		if err := rows.Scan(&message.ID, &message.Body, &message.SentDate, &message.Picture, &message.SenderID, &message.RecipientUserID, &message.RecipientGroupID, &message.ReplyMessageID, &message.ForwardedMessageID); err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return messages, nil
}
