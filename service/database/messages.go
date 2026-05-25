package database

import "wasatext/service/structs"

/*
 * ADD OPERATIONS
 */

func (db *appdbimpl) AddMessage(body string, sentDate string, picture string, senderID int, conversationID int) (int64, error) {
	messageID, err := db.GetNewMessageID(conversationID)
	if err != nil {
		return -1, err
	}

	query := "INSERT INTO messages (id, conversation_id, body, sent_date, sender_id) VALUES (?, ?, ?, ?, ?)"
	if _, err := db.c.Exec(query, messageID, conversationID, body, sentDate, senderID); err != nil {
		return -1, err
	}

	// Add the picture of the message if it is present
	if picture != "" {
		query := "UPDATE messages SET picture=? WHERE id=?"
		_, err = db.c.Exec(query, picture, messageID)
		if err != nil {
			return -1, err
		}
	}

	return messageID, nil
}

func (db *appdbimpl) AddForwardedMessage(originalMessageID int, originalConversationID int, senderID int, conversationID int, sentDate string) (int64, error) {
	// Get the body and picture of the message to forward to
	var body, picture string

	query1 := `SELECT body, COALESCE(picture, "") FROM messages WHERE id=? AND conversation_id=?`
	err := db.c.QueryRow(query1, originalMessageID, originalConversationID).Scan(&body, &picture)
	if err != nil {
		return -1, err
	}

	// Create a new message with the same body and picture of the message to forward to
	messageID, err := db.GetNewMessageID(conversationID)
	if err != nil {
		return -1, err
	}

	query2 := "INSERT INTO messages (id, conversation_id, body, sent_date, sender_id, forwarded_message_id) VALUES (?, ?, ?, ?, ?, ?)"
	if _, err := db.c.Exec(query2, messageID, conversationID, body, sentDate, senderID, originalMessageID); err != nil {
		return -1, err
	}

	// Add the picture of the message if it is present
	if picture != "" {
		query := "UPDATE messages SET picture=? WHERE id=?"
		_, err = db.c.Exec(query, picture, messageID)
		if err != nil {
			return -1, err
		}
	}

	return messageID, nil
}

func (db *appdbimpl) AddReplyMessage(body string, sentDate string, picture string, senderID int, conversationID int, replyMessageID int) (int64, error) {
	messageID, err := db.GetNewMessageID(conversationID)
	if err != nil {
		return -1, err
	}

	query := "INSERT INTO messages (id, conversation_id, body, sent_date, sender_id, reply_message_id) VALUES (?, ?, ?, ?, ?, ?)"
	if _, err := db.c.Exec(query, messageID, conversationID, body, sentDate, senderID, replyMessageID); err != nil {
		return -1, err
	}

	if picture != "" {
		query := "UPDATE messages SET picture=? WHERE id=?"
		_, err = db.c.Exec(query, picture, messageID)
		if err != nil {
			return -1, err
		}
	}

	return messageID, nil
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

	query1 := `SELECT id, conversation_id, body, sent_date, COALESCE(picture, ""), sender_id, COALESCE(reply_message_id, -1), COALESCE(forwarded_message_id, -1)
		FROM messages
		WHERE id=? AND conversation_id=?`
	err := db.c.QueryRow(query1, messageID, conversationID).Scan(&message.ID, &message.SentDate, &message.Picture, &message.SenderID, &message.ConversationID, &message.ReplyMessageID, &message.ForwardedMessageID)
	return message, err
}

func (db *appdbimpl) GetMessagesByConversation(conversationID int) ([]structs.Message, error) {
	query := `SELECT id, body, sent_date, COALESCE(picture, ""), sender_id, conversation_id, COALESCE(reply_message_id, -1), COALESCE(forwarded_message_id, -1)
		FROM messages
		WHERE conversation_id=?
		ORDER BY sent_date DESC`
	rows, err := db.c.Query(query, conversationID)
	if err != nil {
		return nil, err
	}

	var messages []structs.Message
	for rows.Next() {
		var message structs.Message
		if err := rows.Scan(&message.ID, &message.ConversationID, &message.Body, &message.SentDate, &message.Picture, &message.SenderID, &message.ReplyMessageID, &message.ForwardedMessageID); err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return messages, nil
}

func (db *appdbimpl) GetLastMessage(conversationID int) (structs.Message, error) {
	var message structs.Message

	query := `SELECT id, conversation_id, body, sent_date, COALESCE(picture, ""), sender_id, COALESCE(reply_message_id, -1), COALESCE(forwarded_message_id, -1)
		FROM messages
		WHERE conversation_id=?
		ORDER BY sent_date DESC LIMIT 1`
	err := db.c.QueryRow(query, conversationID).Scan(&message.ID, &message.ConversationID, &message.SentDate, &message.Picture, &message.SenderID, &message.ReplyMessageID, &message.ForwardedMessageID)
	return message, err
}
