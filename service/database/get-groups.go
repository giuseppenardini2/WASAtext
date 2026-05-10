package database

import "wasatext/service/structs"

// Given a group ID, return the group's members
func (db *appdbimpl) GetGroups(userID int) ([]structs.Group, error) {
	query := "SELECT g.id, g.name, g.picture, g.user_creator_id FROM groups AS g, users_groups AS ug WHERE g.id=ug.group_id AND ug.user_id=?"
	rows, err := db.c.Query(query, userID)
	if err != nil {
		return nil, err
	}

	var groups []structs.Group
	for rows.Next() {
		var group structs.Group
		if err := rows.Scan(&group.Id, &group.Name, &group.Picture, &group.UserCreatorID); err != nil {
			return nil, err
		}
		groups = append(groups, group)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return groups, nil
}
