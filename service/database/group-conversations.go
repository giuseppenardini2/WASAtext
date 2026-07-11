package database

import "wasatext/service/structs"

/*
 * ADD OPERATIONS
 */

func (db *appdbimpl) AddGroupConversation(name string, pictureURI string, userCreatorID int) (structs.GroupConversation, error) {
	tx, err := db.c.Begin()
	if err != nil {
		return structs.GroupConversation{}, err
	}
	defer func() {
		if err := tx.Rollback(); err != nil {
			return
		}
	}()

	query1 := "INSERT INTO conversations DEFAULT VALUES"
	result, err := tx.Exec(query1)
	if err != nil {
		return structs.GroupConversation{}, err
	}

	groupConversationID, err := result.LastInsertId()
	if err != nil {
		return structs.GroupConversation{}, err
	}

	query2 := "INSERT INTO group_conversations (id, name, user_creator_id) VALUES (?, ?, ?)"
	if _, err := tx.Exec(query2, groupConversationID, name, userCreatorID); err != nil {
		return structs.GroupConversation{}, err
	}

	query3 := "INSERT INTO group_conversations_members (user_id, group_conversation_id) VALUES (?, ?)"
	if _, err := tx.Exec(query3, userCreatorID, groupConversationID); err != nil {
		return structs.GroupConversation{}, err
	}

	if pictureURI != "" {
		query := "UPDATE group_conversations SET picture_uri=? WHERE id=?"
		if _, err := tx.Exec(query, pictureURI, groupConversationID); err != nil {
			return structs.GroupConversation{}, err
		}
	}

	if err := tx.Commit(); err != nil {
		return structs.GroupConversation{}, err
	}

	return structs.GroupConversation{
		ID:            int(groupConversationID),
		Name:          name,
		PictureURI:    pictureURI,
		UserCreatorID: userCreatorID,
	}, nil
}

/*
 * UPDATE OPERATIONS
 */

func (db *appdbimpl) UpdateGroupConversationName(groupConversationID int, name string) error {
	query := "UPDATE group_conversations SET name=? WHERE id=?"
	_, err := db.c.Exec(query, name, groupConversationID)
	return err
}

func (db *appdbimpl) UpdateGroupConversationPicture(groupConversationID int, pictureURI string) error {
	query := "UPDATE group_conversations SET picture_uri=? WHERE id=?"
	_, err := db.c.Exec(query, pictureURI, groupConversationID)
	return err
}

/*
 * DELETE OPERATIONS
 */

func (db *appdbimpl) DeleteGroupConversation(groupConversationID int) error {
	query := "DELETE FROM group_conversations WHERE id=?"
	_, err := db.c.Exec(query, groupConversationID)
	return err
}

/*
 * GET OPERATIONS
 */

func (db *appdbimpl) GetGroupConversation(groupConversationID int) (structs.GroupConversation, error) {
	var group structs.GroupConversation

	query := "SELECT id, name, COALESCE(picture_uri, ''), user_creator_id FROM group_conversations WHERE id=?"
	err := db.c.QueryRow(query, groupConversationID).Scan(&group.ID, &group.Name, &group.PictureURI, &group.UserCreatorID)
	return group, err
}

func (db *appdbimpl) GetGroupConversationPictureURI(groupConversationID int) (string, error) {
	var pictureURI string

	query := "SELECT COALESCE(picture_uri, '') FROM group_conversations WHERE id=?"
	err := db.c.QueryRow(query, groupConversationID).Scan(&pictureURI)
	return pictureURI, err
}

/*
 * CHECK OPERATIONS
 */

func (db *appdbimpl) CheckGroupConversation(groupConversationID int) error {
	var count int

	query := "SELECT COUNT(id) FROM group_conversations WHERE id=?"
	err := db.c.QueryRow(query, groupConversationID).Scan(&count)
	if err != nil {
		return err
	}
	if count == 0 {
		return ErrNonexistentResource
	}
	return nil
}
