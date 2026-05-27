package database

import "wasatext/service/structs"

/*
 * ADD OPERATIONS
 */

func (db *appdbimpl) AddReaction(userID int, conversationID int, messageID int, emoji string) error {
	query := "INSERT INTO reactions (user_id, conversation_id, message_id, emoji) VALUES (?, ?, ?, ?)"
	_, err := db.c.Exec(query, userID, conversationID, messageID, emoji)
	return err
}

/*
 * DELETE OPERATIONS
 */

func (db *appdbimpl) DeleteReaction(userID int, conversationID int, messageID int) error {
	query := "DELETE FROM reactions WHERE user_id=? AND conversation_id=? AND message_id=?"
	_, err := db.c.Exec(query, userID, conversationID, messageID)
	return err
}

/*
 * GET OPERATIONS
 */

func (db *appdbimpl) GetReaction(userID int, conversationID int, messageID int) (structs.Reaction, error) {
	var reaction structs.Reaction

	query := `SELECT user_id, conversation_id, message_id, emoji
		FROM reactions
		WHERE user_id=? AND conversation_id=? AND message_id=?`
	err := db.c.QueryRow(query, userID, conversationID, messageID).Scan(&reaction.UserID, &reaction.ConversationID, &reaction.MessageID, &reaction.Emoji)
	return reaction, err
}

func (db *appdbimpl) GetReactions(conversationID int, messageID int) ([]structs.Reaction, error) {
	query := `SELECT user_id, conversation_id, message_id, emoji
		FROM reactions
		WHERE conversation_id=? AND message_id=?`
	rows, err := db.c.Query(query, conversationID, messageID)
	if err != nil {
		return nil, err
	}

	var reactions []structs.Reaction
	for rows.Next() {
		var reaction structs.Reaction
		if err := rows.Scan(&reaction.UserID, &reaction.ConversationID, &reaction.MessageID, &reaction.Emoji); err != nil {
			return nil, err
		}
		reactions = append(reactions, reaction)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return reactions, nil
}
