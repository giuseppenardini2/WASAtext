package database

import "wasatext/service/api/structs"

// Given a group ID, return the group's members
func (db *appdbimpl) GetMembersFromGroup(id int) (structs.User[], error) {
	rows, err := db.c.Query("SELECT u.id, u.username, u.picture FROM User u, user_group ug WHERE u.id=ug.user AND ug.group=?", id)
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