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
	"fmt"
)

// AppDatabase is the high level interface for the DB
type AppDatabase interface {
	/*GetName() (string, error)
	SetName(name string) error*/
	

	Ping() error
}

type appdbimpl struct {
	c *sql.DB
}

// New returns a new instance of AppDatabase based on the SQLite connection `db`.
// `db` is required - an error will be returned if `db` is `nil`.
func New(db *sql.DB) (AppDatabase, error) {
	
	//Check if the database is nil
	if db == nil {
		return nil, errors.New("database is required when building a AppDatabase")
	}

	//Enable foreign keys
	sqlForeignKeys := 'PRAGMA foreign_keys = ON'
	if _, err := db.Exec(sqlForeignKeys); err != nil {
		return err
	}
	
	//Tables creation (if they not exists)
	//TODO: Inserire ON DELETE CASCADE ??
	userQuery := 'CREATE TABLE IF NOT EXISTS User (
		id INTEGER PRIMARY KEY,
		username TEXT NOT NULL UNIQUE,
		picture BLOB
	);'
	if _, err := db.Exec(userQuery); err != nil {
		return err
	}

	groupQuery := 'CREATE TABLE IF NOT EXISTS Group (
		id INTEGER PRIMARY KEY,
		name TEXT NOT NULL,
		picture BLOB,
		user_creator INTEGER NOT NULL REFERENCES User(id)
	);'
	if _, err := db.Exec(groupQuery); err != nil {
		return err
	}

	userGroupQuery := 'CREATE TABLE IF NOT EXISTS user_group (
		user INTEGER REFERENCES User(id),
		group INTEGER REFERENCES Group(id),
		PRIMARY KEY(user, group)
	);'
	if _, err := db.Exec(userGroupQuery); err != nil {
		return err
	}

	messageQuery := 'CREATE TABLE IF NOT EXISTS Message (
		id INTEGER PRIMARY KEY,
		body TEXT NOT NULL,
		date TEXT NOT NULL,
		picture BLOB,
		sender INTEGER NOT NULL REFERENCES User(id),
		recipient_user INTEGER REFERENCES User(id),
		recipient_group INTEGER REFERENCES Group(id),
		reply_message INTEGER REFERENCES Message(id),
		forwarded_message INTEGER REFERENCES Message(id),
		CHECK((recipient_user IS NOT NULL AND recipient_group IS NULL) OR
			(recipient_user IS NULL AND recipient_group IS NOT NULL))
	);'
	if _, err := db.Exec(messageQuery); err != nil {
		return err
	}

	userReadsMessageQuery := 'CREATE TABLE IF NOT EXISTS user_reads_message (
		user INTEGER REFERENCES User(id),
		message INTEGER REFERENCES Message(id),
		PRIMARY KEY(user, message)
	);'
	if _, err := db.Exec(userReadsMessageQuery); err != nil {
		return err
	}

	userReactsMessageQuery := 'CREATE TABLE IF NOT EXISTS user_reacts_message (
		user INTEGER REFERENCES User(id),
		message INTEGER REFERENCES Message(id),
		reaction TEXT NOT NULL,
		PRIMARY KEY(user, message)
	);'
	if _, err := db.Exec(userReactsMessageQuery); err != nil {
		return err
	}

	return &appdbimpl{
		c: db,
	}, nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
