package database

import "wasatext/service/structs"

/*
 * ADD OPERATIONS
 */
func (db *appdbimpl) AddGroupConversation(name string, picture string, userCreatorID int) (int64, error) {
	groupConversationID, err := db.AddConversation() //2DO: Come gestire gli id
	if err != nil {
		return -1, err
	}

	query := "INSERT INTO group_conversations (id, name, picture, user_creator_id) VALUES (?, ?, ?)"
	_, err = db.c.Exec(query, groupConversationID, name, picture, userCreatorID)
	if err != nil {
		return -1, err
	}
	return groupConversationID, nil
}

/*
 * UPDATE OPERATIONS
 */
// Given a group ID and a profile picture, update the profile picture of the group
func (db *appdbimpl) UpdateGroupConversationPicture(groupConversationID int, picture string) error {
	query := "UPDATE group_conversations SET picture=? WHERE id=?"
	_, err := db.c.Exec(query, picture, groupConversationID)
	return err
}

/*
 * DELETE OPERATIONS
 */
// Given a group ID, delete the group
func (db *appdbimpl) DeleteGroupConversation(groupConversationID int) error {
	query := "DELETE FROM group_conversations WHERE id=?"
	_, err := db.c.Exec(query, groupConversationID)
	return err
}

/*
 * GET OPERATIONS
 */
// Given a group ID, return the group
func (db *appdbimpl) GetGroupConversation(groupConversationID int) (structs.Group, error) {
	var group structs.Group

	query := "SELECT id, name, picture, user_creator_id FROM group_conversations WHERE id=?"
	err := db.c.QueryRow(query, groupConversationID).Scan(&group.ID, &group.Name, &group.Picture, &group.UserCreatorID)
	return group, err
}
