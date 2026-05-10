package database

import "wasatext/service/structs"

// Given a user ID, return the user's messages as conversation
func (db *appdbimpl) GetMessages(userID int) ([]structs.Conversation, error) {
	var conversations []structs.Conversation

	//TODO: Questa query prende solo i messaggi inviati, deve prendere anche quelli ricevuti
	query1 := `SELECT DISTINCT u.id, u.username, u.picture, (
			SELECT body
			FROM messages
			WHERE (sender_id=? AND recipient_user_id=u.id) OR (sender_id=u.id AND recipient_user_id=?)
			ORDER BY sent_date DESC LIMIT 1
		) message_body, (
		 	SELECT picture
			FROM messages
			WHERE (sender_id=? AND recipient_user_id=u.id) OR (sender_id=u.id AND recipient_user_id=?)
			ORDER BY sent_date DESC LIMIT 1
		) message_picture, (
		 	SELECT sent_date
			FROM messages
			WHERE (sender_id=? AND recipient_user_id=u.id) OR (sender_id=u.id AND recipient_user_id=?)
			ORDER BY sent_date DESC LIMIT 1
		) message_sent_date
		FROM messages AS m, users AS u
		WHERE (m.sender_id=? AND m.recipient_user_id=u.id) OR (m.sender_id=u.id AND m.recipient_user_id=?)`
	rows, err := db.c.Query(query1, userID, userID, userID, userID, userID, userID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var conversation structs.Conversation
		conversation.ConversationType = structs.UserConversation
		if err := rows.Scan(&conversation.RecipientID, &conversation.RecipientName, &conversation.RecipientPicture, &conversation.LastMessageBody, &conversation.LastMessagePicture, &conversation.LastMessageSentDate); err != nil {
			return nil, err
		}
		conversations = append(conversations, conversation)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	query2 := `SELECT DISTINCT g.id, g.name, g.picture, (
			SELECT m2.body
			FROM messages AS m2, users_groups AS ug2
			WHERE ug2.user_id=? AND ug2.group_id=g.id AND m2.recipient_group_id=g.id
			ORDER BY m2.sent_date DESC LIMIT 1
		) message_body, (
		 	SELECT m2.picture
			FROM messages AS m2, users_groups AS ug2
			WHERE ug2.user_id=? AND ug2.group_id=g.id AND m2.recipient_group_id=g.id
			ORDER BY m2.sent_date DESC LIMIT 1
		) message_picture, (
		 	SELECT m2.sent_date
			FROM messages AS m2, users_groups AS ug2
			WHERE ug2.user_id=? AND ug2.group_id=g.id AND m2.recipient_group_id=g.id
			ORDER BY m2.sent_date DESC LIMIT 1
		) message_sent_date
	FROM messages AS m1, groups AS g, users_groups AS ug1
	WHERE ug1.user_id=? AND ug1.group_id=g.id AND m1.recipient_group_id=g.id`
	rows, err = db.c.Query(query2, userID, userID, userID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var conversation structs.Conversation
		conversation.ConversationType = structs.GroupConversation
		if err := rows.Scan(&conversation.RecipientID, &conversation.RecipientName, &conversation.RecipientPicture, &conversation.LastMessageBody, &conversation.LastMessagePicture, &conversation.LastMessageSentDate); err != nil {
			return nil, err
		}
		conversations = append(conversations, conversation)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return conversations, nil
}
