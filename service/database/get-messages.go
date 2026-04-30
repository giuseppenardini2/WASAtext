package database

import "wasatext/service/api/structs"

// Given a user ID, return the user's messages as conversation
func (db *appdbimpl) GetMessages(user_id int) ([]structs.Conversation, error) {
	var conversations []structs.Conversation

	rows, err := db.c.Query("SELECT m.recipient_user, u.nickanme, u.picture, (
			SELECT body FROM Message WHERE sender=? AND recipient_user=u.id
		) message_body
	FROM Message m, User u
	WHERE sender=? AND recipient_user IS NOT NULL AND m.recipient_user=u.id", user_id, user_id)
	if err != nil {
		return nil, err
	}
	for rows.Next(){
		var conversation structs.Conversation
		conversation.conversation_type=UserConversation
		if err := rows.Scan(&conversation.recipient_id, &conversation.name, &conversation.picture, &conversation.last_message); err != nil {
			return nil, err
		}
		conversations = append(conversations, conversation)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	rows, err := db.c.Query("SELECT m.recipient_group, u.nickanme, u.picture, (
			SELECT body FROM Message WHERE sender=? AND recipient_group=u.id
		) message_body
	FROM Message m, User u
	WHERE sender=? AND recipient_group IS NOT NULL AND m.recipient_group=u.id", user_id, user_id)
	if err != nil {
		return nil, err
	}
	for rows.Next(){
		var conversation structs.Conversation
		conversation.conversation_type=GroupConversation
		if err := rows.Scan(&conversation.recipient_id, &conversation.name, &conversation.picture, &conversation.last_message); err != nil {
			return nil, err
		}
		conversations = append(conversations, conversation)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return conversations, nil
}