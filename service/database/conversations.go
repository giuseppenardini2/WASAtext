package database

import (
	"errors"
	"wasatext/service/structs"
)

/*
 * ADD OPERATIONS
 */

func (db *appdbimpl) AddConversation() (int64, error) {
	query := "INSERT INTO conversations DEFAULT VALUES"
	result, err := db.c.Exec(query)
	if err != nil {
		return -1, err
	}
	return result.LastInsertId()
}

/*
 * GET OPERATIONS
 */

func (db *appdbimpl) GetConversations(userID int) ([]structs.ConversationPreview, error) {
	var conversations []structs.ConversationPreview

	query1 := `SELECT DISTINCT pc.id, u.username, COALESCE(u.picture, ""), (
			SELECT body
			FROM messages
			WHERE conversation_id=pc.id
			ORDER BY sent_date DESC LIMIT 1
		) message_body, (
		 	SELECT COALESCE(picture, "")
			FROM messages
			WHERE conversation_id=pc.id
			ORDER BY sent_date DESC LIMIT 1
		) message_picture, (
		 	SELECT sent_date
			FROM messages
			WHERE conversation_id=pc.id
			ORDER BY sent_date DESC LIMIT 1
		) message_sent_date
		FROM users AS u, private_conversations AS pc
		WHERE (pc.first_user_id=? OR pc.second_user_id=?) AND u.id!=? AND (pc.first_user_id=u.id OR pc.second_user_id=u.id)`
	rows, err := db.c.Query(query1, userID, userID, userID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var conversationPreview structs.ConversationPreview
		conversationPreview.Type = structs.PrivateConversationType
		if err := rows.Scan(&conversationPreview.ID, &conversationPreview.RecipientName, &conversationPreview.RecipientPicture, &conversationPreview.LastMessageBody, &conversationPreview.LastMessagePicture, &conversationPreview.LastMessageSentDate); err != nil {
			return nil, err
		}
		conversations = append(conversations, conversationPreview)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	query2 := `SELECT DISTINCT gc.id, gc.name, COALESCE(gc.picture, ""), (
			SELECT body
			FROM messages
			WHERE conversation_id=gc.id
			ORDER BY sent_date DESC LIMIT 1
		) message_body, (
		 	SELECT COALESCE(picture, "")
			FROM messages
			WHERE conversation_id=gc.id
			ORDER BY sent_date DESC LIMIT 1
		) message_picture, (
		 	SELECT sent_date
			FROM messages
			WHERE conversation_id=gc.id
			ORDER BY sent_date DESC LIMIT 1
		) message_sent_date
	FROM groups_conversations AS gc, group_conversations_members AS gcm
	WHERE gcm.user_id=? AND gcm.group_id=gc.id`
	rows, err = db.c.Query(query2, userID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var conversationPreview structs.ConversationPreview
		conversationPreview.Type = structs.GroupConversationType
		if err := rows.Scan(&conversationPreview.ID, &conversationPreview.RecipientName, &conversationPreview.RecipientPicture, &conversationPreview.LastMessageBody, &conversationPreview.LastMessagePicture, &conversationPreview.LastMessageSentDate); err != nil {
			return nil, err
		}
		conversations = append(conversations, conversationPreview)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return conversations, nil
}

func (db *appdbimpl) GetConversationType(conversationID int) (structs.ConversationType, error) {
	var count int

	query1 := "SELECT COUNT(id) FROM private_conversations WHERE id=?"
	if err := db.c.QueryRow(query1, conversationID).Scan(&count); err != nil {
		return structs.PrivateConversationType, err
	}
	if count != 0 {
		return structs.PrivateConversationType, nil
	}

	query2 := "SELECT COUNT(id) FROM group_conversations WHERE id=?"
	if err := db.c.QueryRow(query2, conversationID).Scan(&count); err != nil {
		return structs.GroupConversationType, err
	}
	if count != 0 {
		return structs.GroupConversationType, nil
	}

	return structs.GroupConversationType, errors.New("Conversation type error")
}

/*
 * CHECK OPERATIONS
 */

func (db *appdbimpl) CheckUserConversation(userID int, conversationID int) error {
	conversationType, err := db.GetConversationType(conversationID)
	if err != nil {
		return err
	}

	var count int
	switch conversationType {
	case structs.PrivateConversationType:
		query := `SELECT COUNT(user_id) 
		FROM private_conversations 
		WHERE (conversation_id=? AND (first_user_id=? OR second_user_id=?)`
		if err := db.c.QueryRow(query, conversationID, userID, userID).Scan(&count); err != nil {
			return err
		}
	case structs.GroupConversationType:
		query := `SELECT COUNT(user_id) 
			FROM group_conversations_members 
			WHERE user_id=? AND group_conversation_id=?`
		if err := db.c.QueryRow(query, userID, conversationID).Scan(&count); err != nil {
			return err
		}
	}

	if count == 0 {
		return errors.New("User not present in conversation")
	}
	return nil

}
