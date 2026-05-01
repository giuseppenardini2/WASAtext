package database

import "wasatext/service/api/structs"

// Given a user ID, return the user's messages as conversation
func (db *appdbimpl) GetMessages(user_id int) ([]structs.Conversation, error) {
	var conversations []structs.Conversation

	var query1 = `SELECT m.recipient_user, u.nickanme, u.picture, (
			SELECT body FROM messages WHERE sender=? AND recipient_user=u.id
		) message_body
		FROM messages AS m, users AS u
		WHERE sender=? AND recipient_user IS NOT NULL AND m.recipient_user=u.id`
	rows, err := db.c.Query(query1, user_id, user_id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var conversation structs.Conversation
		conversation.Type = structs.UserConversation
		if err := rows.Scan(&conversation.RecipientId, &conversation.Name, &conversation.Picture, &conversation.LastMessage); err != nil {
			return nil, err
		}
		conversations = append(conversations, conversation)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	var query2 = `SELECT m.recipient_group, u.nickanme, u.picture, (
			SELECT body FROM messages WHERE sender=? AND recipient_group=u.id
		) message_body
	FROM messages AS m, users AS u
	WHERE sender=? AND recipient_group IS NOT NULL AND m.recipient_group=u.id`
	rows, err = db.c.Query(query2, user_id, user_id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var conversation structs.Conversation
		conversation.Type = structs.GroupConversation
		if err := rows.Scan(&conversation.RecipientId, &conversation.Name, &conversation.Picture, &conversation.LastMessage); err != nil {
			return nil, err
		}
		conversations = append(conversations, conversation)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return conversations, nil
}
