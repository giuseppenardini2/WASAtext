package database

import "errors"

// Given a user ID and a group, return if the user is a member of the group
func (db *appdbimpl) CheckReadMessage(message_id int) (bool, error) {
	//Check if the recipient of the message is a user or a group
	var recipient_type string
	query1 := `SELECT
		CASE
			WHEN recipient_user_id IS NOT NULL THEN 'recipient_user'
			WHEN recipient_group_id IS NOT NULL THEN 'recipient_group'
			ELSE 'undefined_recipient'
		END recipient_type
		FROM messages
		WHERE id=?`
	err := db.c.QueryRow(query1, message_id).Scan(&recipient_type)
	if err != nil {
		return false, err
	}

	//If the recipient is a user, check if that user read the message. Else if it is a group, check if all group members read the message
	var users_read_message_number int
	query2 := `SELECT COUNT(urm.message_id) 
			FROM users_read_messages
			WHERE urm.message_id=?`
	err = db.c.QueryRow(query2, message_id).Scan(&users_read_message_number)
	if err != nil || users_read_message_number == 0 {
		return false, err
	}
	switch recipient_type {
	case "recipient_user":
		return users_read_message_number != 0, err
	case "recipient_group":
		var group_members_number int
		query3 := `SELECT COUNT(ug.user_id)
			FROM users_groups AS ug, messages AS m
			WHERE m.id=? AND ug.group_id=m.recipient_group_id`
		err = db.c.QueryRow(query3, message_id).Scan(&group_members_number)
		if err != nil {
			return false, err
		}
		return users_read_message_number == group_members_number, err
	default:
		return false, errors.New("An error occurs in the recipient type")
	}
}
