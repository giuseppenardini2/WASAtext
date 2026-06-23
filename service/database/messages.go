package database

import (
	"wasatext/service/structs"
)

/*
 * ADD OPERATIONS
 */

func (db *appdbimpl) AddMessage(body string, pictureURI string, senderID int, conversationID int) (structs.Message, error) {
	messageID, err := db.GetNewMessageID(conversationID)
	if err != nil {
		return structs.Message{}, err
	}

	sentDate := structs.SentDateNow()
	query := "INSERT INTO messages (id, conversation_id, body, sent_date sender_id) VALUES (?, ?, ?, ?, ?)"
	if _, err := db.c.Exec(query, messageID, conversationID, body, sentDate, senderID); err != nil {
		return structs.Message{}, err
	}

	// Add the picture of the message if it is present
	if pictureURI != "" {
		query := "UPDATE messages SET picture_uri=? WHERE id=?"
		_, err = db.c.Exec(query, pictureURI, messageID)
		if err != nil {
			return structs.Message{}, err
		}
	}

	return structs.Message{
		ID:                 int(messageID),
		ConversationID:     conversationID,
		Body:               body,
		SentDate:           sentDate,
		PictureURI:         pictureURI,
		SenderID:           senderID,
		ReplyMessageID:     -1,
		ForwardedMessageID: -1,
	}, nil
}

func (db *appdbimpl) AddReplyMessage(body string, pictureURI string, senderID int, conversationID int, replyMessageID int) (structs.Message, error) {
	messageID, err := db.GetNewMessageID(conversationID)
	if err != nil {
		return structs.Message{}, err
	}

	sentDate := structs.SentDateNow()
	query := "INSERT INTO messages (id, conversation_id, body, sent_date, sender_id, reply_message_id) VALUES (?, ?, ?, ?, ?, ?)"
	if _, err := db.c.Exec(query, messageID, conversationID, body, sentDate, senderID, replyMessageID); err != nil {
		return structs.Message{}, err
	}

	if pictureURI != "" {
		query := "UPDATE messages SET picture_uri=? WHERE id=?"
		_, err = db.c.Exec(query, pictureURI, messageID)
		if err != nil {
			return structs.Message{}, err
		}
	}

	return structs.Message{
		ID:                 int(messageID),
		ConversationID:     conversationID,
		Body:               body,
		SentDate:           sentDate,
		PictureURI:         pictureURI,
		SenderID:           senderID,
		ReplyMessageID:     replyMessageID,
		ForwardedMessageID: -1,
	}, nil
}

func (db *appdbimpl) AddForwardedMessage(originalMessageID int, originalConversationID int, senderID int, conversationID int) (structs.Message, error) {
	// Get the body and pictureURI of the message to forward to
	var body, pictureURI string

	query1 := "SELECT body, COALESCE(picture_uri, '') FROM messages WHERE id=? AND conversation_id=?"
	err := db.c.QueryRow(query1, originalMessageID, originalConversationID).Scan(&body, &pictureURI)
	if err != nil {
		return structs.Message{}, err
	}

	// Create a new message with the same body and picture of the message to forward to
	messageID, err := db.GetNewMessageID(conversationID)
	if err != nil {
		return structs.Message{}, err
	}

	sentDate := structs.SentDateNow()
	query2 := "INSERT INTO messages (id, conversation_id, body, sent_date, sender_id, forwarded_message_id) VALUES (?, ?, ?, ?, ?, ?)"
	if _, err := db.c.Exec(query2, messageID, conversationID, body, sentDate, senderID, originalMessageID); err != nil {
		return structs.Message{}, err
	}

	// Add the picture of the message if it is present
	if pictureURI != "" {
		query := "UPDATE messages SET picture_uri=? WHERE id=?"
		_, err = db.c.Exec(query, pictureURI, messageID)
		if err != nil {
			return structs.Message{}, err
		}
	}

	return structs.Message{
		ID:                 int(messageID),
		ConversationID:     conversationID,
		Body:               body,
		SentDate:           sentDate,
		PictureURI:         pictureURI,
		SenderID:           senderID,
		ReplyMessageID:     -1,
		ForwardedMessageID: originalMessageID,
	}, nil
}

/*
 * DELETE OPERATIONS
 */

func (db *appdbimpl) DeleteMessage(messageID int, conversationID int) error {
	query := "DELETE FROM messages WHERE id=? AND conversation_id=?"
	_, err := db.c.Exec(query, messageID, conversationID)
	return err
}

/*
 * GET OPERATIONS
 */

func (db *appdbimpl) GetNewMessageID(conversationID int) (int64, error) {
	var maxID int64
	query := "SELECT MAX(id) FROM messages WHERE conversation_id=?"
	if err := db.c.QueryRow(query, conversationID).Scan(&maxID); err != nil {
		return -1, err
	}
	return maxID + 1, nil
}

func (db *appdbimpl) GetMessage(conversationID int, messageID int) (structs.Message, error) {
	var message structs.Message

	query := `SELECT id, conversation_id, body, sent_date, COALESCE(picture_uri, X''), sender_id, COALESCE(reply_message_id, -1), COALESCE(forwarded_message_id, -1)
		FROM messages
		WHERE id=? AND conversation_id=?`
	err := db.c.QueryRow(query, messageID, conversationID).Scan(&message.ID, &message.SentDate, &message.PictureURI, &message.SenderID, &message.ConversationID, &message.ReplyMessageID, &message.ForwardedMessageID)
	return message, err
}

func (db *appdbimpl) GetMessages(conversationID int) (structs.Conversation, error) {
	query := `SELECT id, body, sent_date, COALESCE(picture_uri, X''), sender_id, conversation_id, COALESCE(reply_message_id, -1), COALESCE(forwarded_message_id, -1)
		FROM messages
		WHERE conversation_id=?
		ORDER BY sent_date DESC`
	rows, err := db.c.Query(query, conversationID)
	if err != nil {
		return structs.Conversation{}, err
	}

	var messages []structs.Message
	for rows.Next() {
		var message structs.Message
		if err := rows.Scan(&message.ID, &message.ConversationID, &message.Body, &message.SentDate, &message.PictureURI, &message.SenderID, &message.ReplyMessageID, &message.ForwardedMessageID); err != nil {
			return structs.Conversation{}, err
		}
		messages = append(messages, message)
	}
	if err := rows.Err(); err != nil {
		return structs.Conversation{}, err
	}

	return structs.Conversation{
		ID:       conversationID,
		Messages: messages,
	}, nil
}

func (db *appdbimpl) GetLastMessage(conversationID int) (structs.Message, error) {
	var message structs.Message

	query := `SELECT id, conversation_id, body, sent_date, COALESCE(picture_uri, X''), sender_id, COALESCE(reply_message_id, -1), COALESCE(forwarded_message_id, -1)
		FROM messages
		WHERE conversation_id=?
		ORDER BY sent_date DESC LIMIT 1`
	err := db.c.QueryRow(query, conversationID).Scan(&message.ID, &message.ConversationID, &message.SentDate, &message.PictureURI, &message.SenderID, &message.ReplyMessageID, &message.ForwardedMessageID)
	return message, err
}

/*
 * CHECK OPERATIONS
 */

func (db *appdbimpl) CheckMessageConversation(messageID int, conversationID int) error {
	var count int

	query := "SELECT COUNT(id) FROM messages WHERE id=? AND conversation_id=?"
	err := db.c.QueryRow(query, messageID, conversationID).Scan(&count)
	if err != nil {
		return err
	}
	if count == 0 {
		return ErrNonexistentResource
	}
	return nil
}
