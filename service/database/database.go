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
	AddUser(username string, picture structs.Picture) (structs.User, error)
	UpdateUserUsername(userID int, username string) error
	UpdateUserPicture(userID int, picture structs.Picture) error
	GetUser(userID int) (structs.User, error)
	GetUserByUsername(username string) (structs.User, error)
	GetUsers(username string) ([]structs.User, error)
	CheckUser(userID int) error
	CheckUserUsername(username string) error

	AddConversation() (int64, error)
	GetConversations(userID int) ([]structs.ConversationPreview, error)
	GetConversationType(conversationID int) (structs.ConversationType, error)
	CheckUserConversation(userID int, conversationID int) error

	AddPrivateConversation(firstUserID int, secondUserID int) (int64, error)
	CheckUserPrivateConversation(firstUserID int, secondUserID int) error

	AddGroupConversation(name string, picture structs.Picture, userCreatorID int) (structs.GroupConversation, error)
	DeleteGroupConversation(groupConversationID int) error
	GetGroupConversation(groupConversationID int) (structs.GroupConversation, error)
	UpdateGroupConversationName(groupConversationID int, name string) error
	UpdateGroupConversationPicture(groupConversationID int, picture structs.Picture) error
	CheckGroupConversation(groupConversationID int) error

	AddGroupConversationMember(userID int, groupConversationID int) error
	DeleteGroupConversationMember(userID int, groupConversationID int) error
	GetGroupConversartionMembers(groupConversationsID int) ([]structs.User, error)
	CheckGroupConversationMember(userID int, groupConversationID int) error

	AddMessage(body string, picture structs.Picture, senderID int, conversationID int) (structs.Message, error)
	AddReplyMessage(body string, picture structs.Picture, senderID int, conversationID int, replyMessageID int) (structs.Message, error)
	AddForwardedMessage(originalMessageID int, originalConversationID int, senderID int, conversationID int) (structs.Message, error)
	DeleteMessage(messageID int, conversationID int) error
	GetNewMessageID(conversationID int) (int64, error)
	GetMessage(conversationID int, messageID int) (structs.Message, error)
	GetMessages(conversationID int) (structs.Conversation, error)
	GetLastMessage(conversationID int) (structs.Message, error)
	CheckMessageConversation(messageID int, conversationID int) error

	AddReadCheckmark(userID int, messageID int, conversationID int) error
	CheckReadMessage(messageID int, conversationID int) error

	AddReaction(userID int, conversationID int, messageID int, emoji string) (structs.Reaction, error)
	DeleteReaction(reactionID int, conversationID int, messageID int) error
	GetNewReactionID(conversationID int, messageID int) (int64, error)
	GetReaction(reactionID int, conversationID int, messageID int) (structs.Reaction, error)
	GetReactions(conversationID int, messageID int) ([]structs.Reaction, error)
	CheckUserReaction(userID int, reactionID int, messageID int, conversationID int) error

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
	if err := createTables(db); err != nil {
		return nil, err
	}

	return &appdbimpl{
		c: db,
	}, nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
