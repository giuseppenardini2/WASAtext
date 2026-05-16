/*
Package database is the middleware between the app database and the code. All data (de)serialization (save/load) from a
persistent database are handled here. Database specific logic should never escape this package.

To use this package you need to apply migrations to the database if needed/wanted, connect to it (using the database
data source name from config), and then initialize an instance of AppDatabase from the DB connection.

For example, this code adds a parameter in `webapi` executable for the database data source name (add it to the
main.WebAPIConfiguration structure):

	DB struct {
		Filename string `conf:""`
	}

This is an example on how to migrate the DB and connect to it:

	// Start Database
	logger.Println("initializing database support")
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		logger.WithError(err).Error("error opening SQLite DB")
		return fmt.Errorf("opening SQLite: %w", err)
	}
	defer func() {
		logger.Debug("database stopping")
		_ = db.Close()
	}()

Then you can initialize the AppDatabase and pass it to the api package.
*/
package database

import (
	"database/sql"
	"errors"
	"wasatext/service/structs"
)

// AppDatabase is the high level interface for the DB
type AppDatabase interface {
	AddUser(username string) (int64, error)
	GetUser(userID int) (structs.User, error)
	GetUserByUsername(username string) (structs.User, error)
	GetUsers(username string) ([]structs.User, error)
	CheckUserUsername(username string) (bool, error)
	UpdateUserPicture(userID int, picture string) error
	UpdateUserUsername(userID int, username string) error

	AddGroup(name string, picture string, userCreatorID int) (int64, error)
	DeleteGroup(groupID int) error
	GetGroup(groupID int) (structs.Group, error)
	GetGroups(userID int) ([]structs.Group, error)
	GetGroupMembers(groupID int) ([]structs.User, error)
	UpdateGroupPicture(groupID int, picture string) error
	AddMemberToGroup(userID int, groupID int) error
	RemoveUserFromGroup(userID int, groupID int) error
	CheckGroupMember(userID int, groupID int) (bool, error)

	GetMessage(messageID int) (structs.Message, error)
	GetMessages(userID int) ([]structs.Conversation, error)
	GetMessagesByRecipient(userID int, recipientID int, recipientType structs.RecipientType) ([]structs.Message, error)
	GetLastMessage(userID int, recipientID int, recipientType structs.RecipientType) (structs.Message, error) // TODO: Vedere se eliminare (Ultimo messaggio già preso in getMessages)
	AddMessage(body string, sentDate string, picture string, senderID int, recipientID int, recipientType structs.RecipientType) (int64, error)
	AddReplyMessage(body string, sentDate string, picture string, senderID int, recipientID int, recipientType structs.RecipientType, replyMessageID int) (int64, error)
	DeleteMessage(messageID int) error
	AddMessageReadByUser(userID int, messageID int) error
	CheckReadMessage(messageID int) (bool, error)
	AddReaction(userID int, messageID int, reaction string) error
	DeleteReaction(userID int, messageID int) error
	AddForwardedMessage(originalMessageID int, senderID int, recipientID int, recipientType structs.RecipientType, sentDate string) (int64, error)

	Ping() error
}

type appdbimpl struct {
	c *sql.DB
}

// New returns a new instance of AppDatabase based on the SQLite connection `db`.
// `db` is required - an error will be returned if `db` is `nil`.
func New(db *sql.DB) (AppDatabase, error) {

	// Check if the database is nil
	if db == nil {
		return nil, errors.New("database is required when building a AppDatabase")
	}

	// Enable foreign keys
	sqlForeignKeys := "PRAGMA foreign_keys = ON"
	if _, err := db.Exec(sqlForeignKeys); err != nil {
		return nil, err
	}

	// Create tables (if they not exists)
	userQuery := `CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY,
		username TEXT NOT NULL UNIQUE,
		picture BLOB
	);`
	if _, err := db.Exec(userQuery); err != nil {
		return nil, err
	}

	groupQuery := `CREATE TABLE IF NOT EXISTS groups (
		id INTEGER PRIMARY KEY,
		name TEXT NOT NULL,
		picture BLOB,
		user_creator_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE
	);`
	if _, err := db.Exec(groupQuery); err != nil {
		return nil, err
	}

	userGroupQuery := `CREATE TABLE IF NOT EXISTS users_groups (
		user_id INTEGER REFERENCES users(id),
		group_id INTEGER REFERENCES groups(id),
		PRIMARY KEY(user_id, group_id)
	);`
	if _, err := db.Exec(userGroupQuery); err != nil {
		return nil, err
	}

	messageQuery := `CREATE TABLE IF NOT EXISTS messages (
		id INTEGER PRIMARY KEY,
		body TEXT NOT NULL,
		sent_date TEXT DEFAULT CURRENT_TIMESTAMP NOT NULL,
		picture BLOB,
		sender_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
		recipient_user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
		recipient_group_id INTEGER REFERENCES groups(id) ON DELETE CASCADE,
		reply_message_id INTEGER REFERENCES messages(id) ON DELETE CASCADE,
		forwarded_message_id INTEGER REFERENCES messages(id) ON DELETE CASCADE,
		CHECK(
			(recipient_user_id IS NOT NULL AND recipient_group_id IS NULL) 
			OR
			(recipient_user_id IS NULL AND recipient_group_id IS NOT NULL)
		)
	);`
	if _, err := db.Exec(messageQuery); err != nil {
		return nil, err
	}

	userReadsMessageQuery := `CREATE TABLE IF NOT EXISTS users_read_messages (
		user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
		message_id INTEGER REFERENCES messages(id) ON DELETE CASCADE,
		PRIMARY KEY(user_id, message_id)
	);`
	if _, err := db.Exec(userReadsMessageQuery); err != nil {
		return nil, err
	}

	userReactsMessageQuery := `CREATE TABLE IF NOT EXISTS users_react_messages (
		user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
		message_id INTEGER REFERENCES messages(id) ON DELETE CASCADE,
		reaction TEXT NOT NULL,
		PRIMARY KEY(user_id, message_id)
	);`
	if _, err := db.Exec(userReactsMessageQuery); err != nil {
		return nil, err
	}

	return &appdbimpl{
		c: db,
	}, nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
