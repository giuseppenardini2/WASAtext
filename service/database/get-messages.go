package database

import "wasatext/service/api/structs"

// Given a user ID, return the user's messages as conversation
func (db *appdbimpl) GetMessages(user_id int) ([]structs.Conversation, error) {
	var conversations []structs.Conversation

	//TODO: Questa query prende solo i messaggi inviati, deve prendere anche quelli ricevuti
	query1 := `SELECT DISTINCT m.recipient_user_id, u.username, u.picture, (
			SELECT body
			FROM messages
			WHERE recipient_user_id=u.id
			ORDER BY sent_date DESC LIMIT 1
		) message_body
		FROM messages AS m, users AS u
		WHERE sender_id=? AND recipient_user_id IS NOT NULL AND m.recipient_user_id=u.id` //TODO: Modificare per includere anche messaggi ricevuti e il sent_date
	rows, err := db.c.Query(query1, user_id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var conversation structs.Conversation
		conversation.Type = structs.UserConversation
		if err := rows.Scan(&conversation.RecipientID, &conversation.Name, &conversation.Picture, &conversation.LastMessage); err != nil {
			return nil, err
		}
		conversations = append(conversations, conversation)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	query2 := `SELECT DISTINCT m.recipient_group_id, g.name, g.picture, (
			SELECT body 
			FROM messages
			WHERE recipient_group_id=g.id
			ORDER BY sent_date DESC
		) message_body
	FROM messages AS m, groups AS g
	WHERE sender_id=? AND recipient_group_id IS NOT NULL AND m.recipient_group_id=g.id`
	rows, err = db.c.Query(query2, user_id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var conversation structs.Conversation
		conversation.Type = structs.GroupConversation
		if err := rows.Scan(&conversation.RecipientID, &conversation.Name, &conversation.Picture, &conversation.LastMessage); err != nil {
			return nil, err
		}
		conversations = append(conversations, conversation)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return conversations, nil
}
