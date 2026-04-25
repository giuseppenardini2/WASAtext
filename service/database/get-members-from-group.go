package database

import "WASAtext/service/api/structs"

// Given a group ID, return the group's members
func (db *appdbimpl) GetMembersFromGroup(id int) (structs.User[], error) {
	rows, err := db.c.Query("SELECT User.id, User.username, User.picture FROM User, user_group WHERE User.id=user_group.user AND user_group.group=?", id)
	if err != nil {
		return nil, err
	}

	var users structs.Users[]
	for rows.Next(){
		var user structs.User
		if err := rows.Scan(&user.id, &user.username, &user.picture); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return users, err
}