package database

import "wasatext/service/api/structs"

// Given a group ID, return the group's members
func (db *appdbimpl) GetGroups(user_id int) ([]structs.Group, error) {
	var query = "SELECT g.id, g.name, g.picture, g.user_creator FROM Group g, user_group ug WHERE g.id=ug.group AND ug.user=?"
	rows, err := db.c.Query(query, user_id)
	if err != nil {
		return nil, err
	}

	var groups []structs.Group
	for rows.Next() {
		var group structs.Group
		if err := rows.Scan(&group.Id, &group.Name, &group.Picture, &group.UserCreator); err != nil {
			return nil, err
		}
		groups = append(groups, group)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return groups, nil
}
