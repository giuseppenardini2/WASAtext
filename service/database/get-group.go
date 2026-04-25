package database

import "WASAtext/service/api/structs"

// Given a group ID, return the group
func (db *appdbimpl) GetGroup(id int) (structs.Group, error) {
	var group structs.Group
	err := db.c.QueryRow("SELECT id, name, picture, user_creator FROM Group WHERE id=?", id).Scan(&group.id, &group.name, &group.picture, &group.user_creator)
	if err != nil {
		return nil, err
	}
	return group, err
}