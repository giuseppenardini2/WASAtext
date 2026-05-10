package database

import "wasatext/service/api/structs"

// Given a group ID, return the group
func (db *appdbimpl) GetGroup(groupID int) (structs.Group, error) {
	var group structs.Group

	query := "SELECT id, name, picture, user_creator_id FROM groups WHERE id=?"
	err := db.c.QueryRow(query, groupID).Scan(&group.Id, &group.Name, &group.Picture, &group.UserCreatorID)
	return group, err
}
