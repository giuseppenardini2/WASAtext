package database

import "WASAtext/service/api/structs"

// Given a user ID, return the user's messages
//TODO: Decidere come prendere le conversazioni
func (db *appdbimpl) GetMessages(user_id int) (structs.Conversation[], error) {
	rows, err := db.c.Query("SELECT m.recipient_user, u.nickanme, u.picture, (
			SELECT body FROM Message WHERE sender=? AND recipient_user=u.id
		) message_body
	FROM Message m, User u
	WHERE sender=? AND recipient_user IS NOT NULL AND m.recipient_user=u.id", user_id, user_id)
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