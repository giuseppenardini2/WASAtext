package database

import "wasatext/service/api/structs"

// Given a user ID and recipient user ID, return the user's messages with that user
func (db *appdbimpl) GetMessagesByRecipientUser(user_id int, recipient_user int) ([]structs.Message, error) {
	var query = `SELECT id, body, date, picture, sender, recipient_user, recipient_group, reply_message, forwarded_message
		FROM messages WHERE sender=? AND recipient_user=?
		ORDER BY date DESC`
	rows, err := db.c.Query(query, user_id, recipient_user)
	if err != nil {
		return nil, err
	}

	var messages []structs.Message
	for rows.Next() {
		var message structs.Message
		if err := rows.Scan(&message.Id, &message.Body, &message.Date, &message.Picture, &message.Sender, &message.RecipientUser, &message.RecipientGroup, &message.ReplyMessage, &message.ForwardedMessage); err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return messages, err
}
