package database

import "WASAtext/service/api/structs"

// Given a group ID, delete the group
func (db *appdbimpl) DeleteGroup(id int) (error) {
	_, err := db.c.Exec("DELETE FROM Group WHERE id=?", id)
	return err
}