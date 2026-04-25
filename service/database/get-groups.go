package database

import "WASAtext/service/api/structs"

// Given a group ID, return the group's members
func (db *appdbimpl) GetGroups(id int) (structs.Group[], error) {
	rows, err := db.c.Query("SELECT Group.id, Group.name, Group.picture, Group.user_creator FROM Group, user_group WHERE Group.id=user_group.group AND user_group.user=?", id)
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
	return groups, err
}