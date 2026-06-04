package database

import "wasatext/service/structs"

/*
 * ADD OPERATIONS
 */

func (db *appdbimpl) AddReaction(userID int, conversationID int, messageID int, emoji string) (structs.Reaction, error) {
	reactionID, err := db.GetNewReactionID(conversationID, messageID)
	if err != nil {
		return structs.Reaction{}, err
	}

	query := "INSERT INTO reactions (id, conversation_id, message_id, user_id, emoji) VALUES (?, ?, ?, ?)"
	_, err = db.c.Exec(query, reactionID, conversationID, messageID, userID, emoji)
	return structs.Reaction{
		ID:             int(reactionID),
		UserID:         userID,
		ConversationID: conversationID,
		MessageID:      messageID,
		Emoji:          emoji,
	}, err
}

/*
 * DELETE OPERATIONS
 */

func (db *appdbimpl) DeleteReaction(reactionID int, conversationID int, messageID int) error {
	query := "DELETE FROM reactions WHERE id=? AND conversation_id=? AND message_id=?"
	_, err := db.c.Exec(query, reactionID, conversationID, messageID)
	return err
}

/*
 * GET OPERATIONS
 */

func (db *appdbimpl) GetNewReactionID(conversationID int, messageID int) (int64, error) {
	var maxID int64
	query := "SELECT MAX(id) FROM reactions WHERE conversation_id=? AND message_id=?"
	if err := db.c.QueryRow(query, conversationID).Scan(&maxID); err != nil {
		return -1, err
	}
	return maxID + 1, nil
}

func (db *appdbimpl) GetReaction(reactionID int, conversationID int, messageID int) (structs.Reaction, error) {
	var reaction structs.Reaction

	query := `SELECT id, conversation_id, message_id, user_id, emoji
		FROM reactions
		WHERE reaction_id=? AND conversation_id=? AND message_id=?`
	err := db.c.QueryRow(query, reactionID, conversationID, messageID).Scan(&reactionID, &reaction.ConversationID, &reaction.MessageID, &reaction.UserID, &reaction.Emoji)
	return reaction, err
}

func (db *appdbimpl) GetReactions(conversationID int, messageID int) ([]structs.Reaction, error) {
	query := `SELECT id, conversation_id, message_id, user_id, emoji
		FROM reactions
		WHERE conversation_id=? AND message_id=?`
	rows, err := db.c.Query(query, conversationID, messageID)
	if err != nil {
		return nil, err
	}

	var reactions []structs.Reaction
	for rows.Next() {
		var reaction structs.Reaction
		if err := rows.Scan(&reaction.ID, &reaction.ConversationID, &reaction.MessageID, &reaction.UserID, &reaction.Emoji); err != nil {
			return nil, err
		}
		reactions = append(reactions, reaction)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return reactions, nil
}

/*
 * CHECK OPERATIONS
 */

func (db *appdbimpl) CheckUserReaction(userID int, reactionID int, messageID int, conversationID int) error {
	var count int

	query := "SELECT COUNT(id) FROM reactions WHERE id=? AND conversation_id=? AND message_id=? AND user_id=?"
	err := db.c.QueryRow(query, reactionID, conversationID, messageID, userID).Scan(&count)
	if err != nil {
		return err
	}
	if count == 0 {
		return ErrNonexistentResource
	}
	return nil
}
