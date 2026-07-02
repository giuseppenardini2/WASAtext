package database

import (
	"database/sql"
	"errors"
	"wasatext/service/structs"
)

/*
 * GET OPERATIONS
 */

func (db *appdbimpl) GetConversations(userID int) ([]structs.ConversationPreview, error) {
	var conversations []structs.ConversationPreview

	query1 := `SELECT DISTINCT pc.id, u.id, u.username, COALESCE(u.picture_uri, '')
		FROM users AS u
		INNER JOIN private_conversations AS pc ON (pc.first_user_id=u.id OR pc.second_user_id=u.id)
		WHERE u.id!=? AND (pc.first_user_id=? OR pc.second_user_id=?)`
	rows1, err := db.c.Query(query1, userID, userID, userID)
	if err != nil {
		return nil, err
	}
	defer rows1.Close()
	query2 := `SELECT id, body, COALESCE(picture_uri, ''), sent_date
		FROM messages
		WHERE conversation_id=?
		ORDER BY sent_date DESC LIMIT 1`
	for rows1.Next() {
		var conversationPreview structs.ConversationPreview
		conversationPreview.Type = structs.PrivateConversationType
		if err := rows1.Scan(&conversationPreview.ID, &conversationPreview.RecipientID, &conversationPreview.RecipientName, &conversationPreview.RecipientPictureURI); err != nil {
			return nil, err
		}
		conversationPreview.LastMessageID = -1
		if err := db.c.QueryRow(query2, conversationPreview.ID).Scan(&conversationPreview.LastMessageID, &conversationPreview.LastMessageBody, &conversationPreview.LastMessagePictureURI, &conversationPreview.LastMessageSentDate); !errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
		conversations = append(conversations, conversationPreview)
	}
	if err = rows1.Err(); err != nil {
		return nil, err
	}

	query3 := `SELECT DISTINCT gc.id, gc.id, gc.name, COALESCE(gc.picture_uri, '')
	FROM group_conversations AS gc
	INNER JOIN group_conversations_members AS gcm ON gcm.group_conversation_id=gc.id
	WHERE gcm.user_id=?`
	rows2, err := db.c.Query(query3, userID)
	if err != nil {
		return nil, err
	}
	defer rows2.Close()
	for rows2.Next() {
		var conversationPreview structs.ConversationPreview
		conversationPreview.Type = structs.GroupConversationType
		if err := rows2.Scan(&conversationPreview.ID, &conversationPreview.RecipientID, &conversationPreview.RecipientName, &conversationPreview.RecipientPictureURI, &conversationPreview.LastMessageBody, &conversationPreview.LastMessagePictureURI, &conversationPreview.LastMessageSentDate); err != nil {
			return nil, err
		}
		if err := db.c.QueryRow(query2, conversationPreview.ID).Scan(&conversationPreview.LastMessageBody, &conversationPreview.LastMessagePictureURI, &conversationPreview.LastMessageSentDate); !errors.Is(err, sql.ErrNoRows) {
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
