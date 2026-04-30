package database

import "wasatext/service/api/structs"

// Given a group ID, return the group's members
func (db *appdbimpl) GetMembersFromGroup(id int) ([]structs.User, error) {
	var query = "SELECT u.id, u.username, u.picture FROM User u, user_group ug WHERE u.id=ug.user AND ug.group=?"
	rows, err := db.c.Query(query, id)
	if err != nil {
		return nil, err
	}

	var users []structs.User
	for rows.Next() {
		var user structs.User
		if err := rows.Scan(&user.Id, &user.Username, &user.Picture); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return users, err
}
