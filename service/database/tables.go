package database

import "database/sql"

func createTables(db *sql.DB) error {
	usersQuery := `CREATE TABLE IF NOT EXISTS users (
		id INTEGER,
		username TEXT NOT NULL UNIQUE,
		picture BLOB,
		PRIMARY KEY(id)
	);`
	if _, err := db.Exec(usersQuery); err != nil {
		return err
	}

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
		CONSTRAINT fk_private_conversations
			FOREIGN KEY (id) REFERENCES conversations(id) ON DELETE CASCADE
			FOREIGN KEY (first_user_id) REFERENCES users(id) ON DELETE CASCADE
			FOREIGN KEY (second_user_id) REFERENCES users(id) ON DELETE CASCADE
	);`
	if _, err := db.Exec(privateConversationsQuery); err != nil {
		return err
	}

	groupConversationsQuery := `CREATE TABLE IF NOT EXISTS group_conversations (
		id INTEGER,
		name TEXT NOT NULL,
		picture BLOB,
		user_creator_id INTEGER NOT NULL,
		PRIMARY KEY(id),
		CONSTRAINT fk_group_conversations
			FOREIGN KEY(id) REFERENCES conversations(id) ON DELETE CASCADE
			FOREIGN KEY(user_creator_id) REFERENCES users(id) ON DELETE CASCADE
	);`
	if _, err := db.Exec(groupConversationsQuery); err != nil {
		return err
	}

	usersGroupsConversationsQuery := `CREATE TABLE IF NOT EXISTS group_conversations_members (
		user_id INTEGER,
		group_conversation_id INTEGER,
		PRIMARY KEY(user_id, group_id),
		CONSTRAINT fk_group_conversations_members
			FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCEDE
			FOREIGN KEY(group_conversation_id) REFERENCES group_conversations(id)
	);`
	if _, err := db.Exec(usersGroupsConversationsQuery); err != nil {
		return err
	}

	messagesQuery := `CREATE TABLE IF NOT EXISTS messages (
		id INTEGER,
		conversation_id INTEGER,
		body TEXT NOT NULL,
		sent_date TEXT DEFAULT CURRENT_TIMESTAMP NOT NULL,
		picture BLOB,
		sender_id INTEGER NOT NULL,
		reply_message_id INTEGER,
		forwarded_message_id INTEGER,
		PRIMARY KEY(id, conversation_id),
		CONSTRAINT fk_messages
			FOREIGN KEY(conversation_id) REFERENCES conversations(id) ON DELETE CASCADE
			FOREIGN KEY(sender_id) REFERENCES users(id) ON DELETE CASCADE
			FOREIGN KEY(reply_message_id) REFERENCES messages(id) ON DELETE CASCADE
			FOREIGN KEY(forwarded_message_id) REFERENCES messages(id) ON DELETE CASCADE
	);`
	if _, err := db.Exec(messagesQuery); err != nil {
		return err
	}

	usersReadMessagesQuery := `CREATE TABLE IF NOT EXISTS read_checkmarks (
		user_id INTEGER,
		message_id INTEGER,
		conversation_id INTEGER,
		PRIMARY KEY(user_id, message_id, conversation_id),
		CONSTRAINT fk_read_checkmarks
			FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE
			FOREING KEY(message_id, conversation_id) REFERENCES messages(id, conversation_id) ON DELETE CASCADE
	);`
	if _, err := db.Exec(usersReadMessagesQuery); err != nil {
		return err
	}

	usersReactMessagesQuery := `CREATE TABLE IF NOT EXISTS reactions (
		user_id INTEGER,
		message_id INTEGER,
		conversation_id INTEGER,
		reaction TEXT NOT NULL,
		PRIMARY KEY(user_id, message_id, conversation_id),
		CONSTRAINT fk_reactions
			FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE
			FOREIGN KEY(message_id, conversation_id) REFERENCES messages(id, conversation_id) ON DELETE CASCADE
	);`
	if _, err := db.Exec(usersReactMessagesQuery); err != nil {
		return err
	}

	return nil
}
