package database

import "database/sql"

func createTables(db *sql.DB) error {
	// Users
	usersQuery := `CREATE TABLE IF NOT EXISTS users (
		id INTEGER,
		username TEXT NOT NULL UNIQUE,
		picture_uri TEXT,
		PRIMARY KEY(id)
	);`
	if _, err := db.Exec(usersQuery); err != nil {
		return err
	}

	// Conversations
	conversationsQuery := `CREATE TABLE IF NOT EXISTS conversations (
		id INTEGER,
		PRIMARY KEY(id)
	);`
	if _, err := db.Exec(conversationsQuery); err != nil {
		return err
	}

	privateConversationsQuery := `CREATE TABLE IF NOT EXISTS private_conversations (
		id INTEGER,
		first_user_id INTEGER NOT NULL,
		second_user_id INTEGER NOT NULL,
		PRIMARY KEY(id),
		UNIQUE(first_user_id, second_user_id),
		FOREIGN KEY (id) REFERENCES conversations(id) ON DELETE CASCADE,
		FOREIGN KEY (first_user_id) REFERENCES users(id) ON DELETE CASCADE,
		FOREIGN KEY (second_user_id) REFERENCES users(id) ON DELETE CASCADE
	);`
	if _, err := db.Exec(privateConversationsQuery); err != nil {
		return err
	}

	groupConversationsQuery := `CREATE TABLE IF NOT EXISTS group_conversations (
		id INTEGER,
		name TEXT NOT NULL,
		picture_uri TEXT,
		user_creator_id INTEGER NOT NULL,
		PRIMARY KEY(id),
		FOREIGN KEY(id) REFERENCES conversations(id) ON DELETE CASCADE,
		FOREIGN KEY(user_creator_id) REFERENCES users(id) ON DELETE CASCADE
	);`
	if _, err := db.Exec(groupConversationsQuery); err != nil {
		return err
	}

	groupConversationsMembersQuery := `CREATE TABLE IF NOT EXISTS group_conversations_members (
		user_id INTEGER,
		group_conversation_id INTEGER,
		PRIMARY KEY(user_id, group_conversation_id),
		FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE,
		FOREIGN KEY(group_conversation_id) REFERENCES group_conversations(id) ON DELETE CASCADE
	);`
	if _, err := db.Exec(groupConversationsMembersQuery); err != nil {
		return err
	}

	// Messages
	messagesQuery := `CREATE TABLE IF NOT EXISTS messages (
		id INTEGER,
		conversation_id INTEGER,
		body TEXT NOT NULL,
		sent_date TEXT DEFAULT CURRENT_TIMESTAMP NOT NULL,
		picture_uri TEXT,
		sender_id INTEGER NOT NULL,
		reply_message_id INTEGER,
		forwarded_message_id INTEGER,
		PRIMARY KEY(id, conversation_id),
		FOREIGN KEY(conversation_id) REFERENCES conversations(id) ON DELETE CASCADE,
		FOREIGN KEY(sender_id) REFERENCES users(id) ON DELETE CASCADE,
		FOREIGN KEY(reply_message_id) REFERENCES messages(id) ON DELETE CASCADE,
		FOREIGN KEY(forwarded_message_id) REFERENCES messages(id) ON DELETE CASCADE
	);`
	if _, err := db.Exec(messagesQuery); err != nil {
		return err
	}

	readCheckmarksQuery := `CREATE TABLE IF NOT EXISTS read_checkmarks (
		user_id INTEGER,
		message_id INTEGER,
		conversation_id INTEGER,
		PRIMARY KEY(user_id, message_id, conversation_id),
		FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE,
		FOREIGN KEY(message_id, conversation_id) REFERENCES messages(id, conversation_id) ON DELETE CASCADE
	);`
	if _, err := db.Exec(readCheckmarksQuery); err != nil {
		return err
	}

	reactionsQuery := `CREATE TABLE IF NOT EXISTS reactions (
		id INTEGER,
		conversation_id INTEGER,
		message_id INTEGER,
		user_id INTEGER NOT NULL,
		emoji TEXT NOT NULL,
		PRIMARY KEY(id, message_id, conversation_id),
		UNIQUE(user_id, message_id, conversation_id),
		FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE,
		FOREIGN KEY(message_id, conversation_id) REFERENCES messages(id, conversation_id) ON DELETE CASCADE
	);`
	if _, err := db.Exec(reactionsQuery); err != nil {
		return err
	}

	return nil
}
