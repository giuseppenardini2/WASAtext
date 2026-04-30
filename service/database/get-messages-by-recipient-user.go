package database

import "wasatext/service/api/structs"

// Given a user ID and recipient user ID, return the user's messages with that user
func (db *appdbimpl) GetMessagesByRecipientUser(user_id int, recipient_user int) (structs.Message[], error) {
	rows, err := db.c.Query("SELECT id, body, date, picture, sender, recipient_user, recipient_group, reply_message, forwarded_message
		FROM Message WHERE sender=? AND recipient_user=?
		ORDER BY date DESC", user_id, recipient_user)
	if err != nil {
		return nil, err
	}

	var messages structs.Message[]
	for rows.Next(){
		var message structs.Message
		if err := rows.Scan(&message.id, &message.body, &message.date, &message.picture, &message.sender, &message.recipient_user, &message.recipient_group, &message.reply_message, &message.forwarded_message); err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return messages, err
}