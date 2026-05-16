package database

import "errors"

// Given a user ID and a group, return if the user is a member of the group
func (db *appdbimpl) CheckReadMessage(messageID int) (bool, error) {
	// Check if the recipient of the message is a user or a group
	var recipientType string
	query1 := `SELECT
		CASE
			WHEN recipient_user_id IS NOT NULL THEN 'recipient_user'
			WHEN recipient_group_id IS NOT NULL THEN 'recipient_group'
			ELSE 'undefined_recipient'
		END recipient_type
		FROM messages
		WHERE id=?`
	err := db.c.QueryRow(query1, messageID).Scan(&recipientType)
	if err != nil {
		return false, err
	}

	// If the recipient is a user, check if that user read the message. Else if it is a group, check if all group members read the message
	var usersReadMessageNumber int
	query2 := `SELECT COUNT(urm.message_id) 
			FROM users_read_messages
			WHERE urm.message_id=?`
	err = db.c.QueryRow(query2, messageID).Scan(&usersReadMessageNumber)
	if err != nil || usersReadMessageNumber == 0 {
		return false, err
	}
	switch recipientType {
	case "recipient_user":
		return usersReadMessageNumber != 0, err
	case "recipient_group":
		var groupMembersNumber int
		query3 := `SELECT COUNT(ug.user_id)
			FROM users_groups AS ug, messages AS m
			WHERE m.id=? AND ug.group_id=m.recipient_group_id`
		err = db.c.QueryRow(query3, messageID).Scan(&groupMembersNumber)
		if err != nil {
			return false, err
		}
		return usersReadMessageNumber == groupMembersNumber, err
	default:
		return false, errors.New("An error occurs in the recipient type")
	}
}
