package database

import (
	"errors"
	"wasatext/service/structs"
)

/*
 * GET OPERATIONS
 */

func (db *appdbimpl) GetConversations(userID int) ([]structs.ConversationPreview, error) {
	var conversations []structs.ConversationPreview

	query1 := `SELECT DISTINCT pc.id, u.username, COALESCE(u.picture_uri, ''), (
			SELECT COALESCE(body, '')
			FROM messages
			WHERE conversation_id=pc.id
			ORDER BY sent_date DESC LIMIT 1
		) message_body, (
		 	SELECT COALESCE(picture_uri, '')
			FROM messages
			WHERE conversation_id=pc.id
			ORDER BY sent_date DESC LIMIT 1
		) message_picture_uri, (
		 	SELECT COALESCE(sent_date, '')
			FROM messages
			WHERE conversation_id=pc.id
			ORDER BY sent_date DESC LIMIT 1
		) message_sent_date
		FROM users AS u, private_conversations AS pc
		WHERE (pc.first_user_id=? OR pc.second_user_id=?) AND u.id!=? AND (pc.first_user_id=u.id OR pc.second_user_id=u.id)`
	rows1, err := db.c.Query(query1, userID, userID, userID)
	if err != nil {
		return nil, err
	}
	defer rows1.Close()
	for rows1.Next() {
		var conversationPreview structs.ConversationPreview
		conversationPreview.Type = structs.PrivateConversationType
		if err := rows1.Scan(&conversationPreview.ID, &conversationPreview.RecipientName, &conversationPreview.RecipientPictureURI, &conversationPreview.LastMessageBody, &conversationPreview.LastMessagePictureURI, &conversationPreview.LastMessageSentDate); err != nil {
			return nil, err
		}
		conversations = append(conversations, conversationPreview)
	}
	if err = rows1.Err(); err != nil {
		return nil, err
	}

	query2 := `SELECT DISTINCT gc.id, gc.name, COALESCE(gc.picture_uri, ''), (
			SELECT COALESCE(body, '')
			FROM messages
			WHERE conversation_id=gc.id
			ORDER BY sent_date DESC LIMIT 1
		) message_body, (
		 	SELECT COALESCE(picture_uri, '')
			FROM messages
			WHERE conversation_id=gc.id
			ORDER BY sent_date DESC LIMIT 1
		) message_picture_uri, (
		 	SELECT COALESCE(sent_date, '')
			FROM messages
			WHERE conversation_id=gc.id
			ORDER BY sent_date DESC LIMIT 1
		) message_sent_date
	FROM group_conversations AS gc, group_conversations_members AS gcm
	WHERE gcm.user_id=? AND gcm.group_conversation_id=gc.id`
	rows2, err := db.c.Query(query2, userID)
	if err != nil {
		return nil, err
	}
	defer rows2.Close()
	for rows2.Next() {
		var conversationPreview structs.ConversationPreview
		conversationPreview.Type = structs.GroupConversationType
		if err := rows2.Scan(&conversationPreview.ID, &conversationPreview.RecipientName, &conversationPreview.RecipientPictureURI, &conversationPreview.LastMessageBody, &conversationPreview.LastMessagePictureURI, &conversationPreview.LastMessageSentDate); err != nil {
			return nil, err
		}
		conversations = append(conversations, conversationPreview)
	}
	if err = rows2.Err(); err != nil {
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
        WHERE conversation_id=? AND (first_user_id=? OR second_user_id=?)`
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
		return ErrNonexistentResource
	}
	return nil

}
