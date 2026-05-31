package database

import "wasatext/service/structs"

/*
 * ADD OPERATIONS
 */

func (db *appdbimpl) AddGroupConversation(name string, picture string, userCreatorID int) (structs.GroupConversation, error) {
	groupConversationID, err := db.AddConversation() //2DO: Come gestire gli id
	if err != nil {
		return structs.GroupConversation{}, err
	}

	query := "INSERT INTO group_conversations (id, name, picture, user_creator_id) VALUES (?, ?, ?)"
	_, err = db.c.Exec(query, groupConversationID, name, picture, userCreatorID)
	if err != nil {
		return structs.GroupConversation{}, err
	}

	if picture != "" {
		query := "UPDATE group_conversations SET picture=? WHERE id=?"
		if _, err := db.c.Exec(query, picture, groupConversationID); err != nil {
			return structs.GroupConversation{}, err
		}
	}

	return structs.GroupConversation{
		ID:            int(groupConversationID),
		Name:          name,
		Picture:       picture,
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

func (db *appdbimpl) UpdateGroupConversationPicture(groupConversationID int, picture string) error {
	query := "UPDATE group_conversations SET picture=? WHERE id=?"
	_, err := db.c.Exec(query, picture, groupConversationID)
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

	query := "SELECT id, name, picture, user_creator_id FROM group_conversations WHERE id=?"
	err := db.c.QueryRow(query, groupConversationID).Scan(&group.ID, &group.Name, &group.Picture, &group.UserCreatorID)
	return group, err
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
		return NonexistentResource
	}
	return nil
}
