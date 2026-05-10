package database

import "wasatext/service/api/structs"

// Given a group ID, return the group members
func (db *appdbimpl) GetGroupMembers(groupID int) ([]structs.User, error) {
	query := "SELECT u.id, u.username, u.picture FROM users AS u, users_groups AS ug WHERE u.id=ug.user_id AND ug.group_id=?"
	rows, err := db.c.Query(query, groupID)
	if err != nil {
		return nil, err
	}

	var users []structs.User
	for rows.Next() {
		var user structs.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Picture); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return users, err
}
