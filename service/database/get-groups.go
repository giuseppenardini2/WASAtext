package database

import "WASAtext/service/api/structs"

// Given a group ID, return the group's members
func (db *appdbimpl) GetGroups(user_id int) (structs.Group[], error) {
	rows, err := db.c.Query("SELECT g.id, g.name, g.picture, g.user_creator FROM Group g, user_group ug WHERE g.id=ug.group AND ug.user=?", user_id)
	if err != nil {
		return nil, err
	}

	var groups structs.Group[]
	for rows.Next(){
		var group structs.Group
		if err := rows.Scan(&group.id, &group.name, &group.picture, &group.user_creator); err != nil {
			return nil, err
		}
		groups = append(groups, group)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return groups, nil
}