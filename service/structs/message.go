package structs

import "time"

type MessageStatus int

const (
	SendingMessageStatus MessageStatus = iota
	SentMessageStatus
	ReadMessageStatus
)

func (messageStatus MessageStatus) String() string {
	return [...]string{"SendingMessageStatus", "SentMessageStatus", "ReadMessageStatus"}[messageStatus]
}

// A struct which represent a message of a conversation
type Message struct {
	ID                             int           `json:"id"`                             // Message ID
	ConversationID                 int           `json:"conversationID"`                 // Conversation ID
	Body                           string        `json:"body"`                           // Message body
	SentDate                       string        `json:"sentDate"`                       // Message sent time
	PictureURI                     string        `json:"pictureURI"`                     // Message picture URI
	Status                         MessageStatus `json:"status"`                         // Status of the message
	SenderID                       int           `json:"senderID"`                       // Sender ID
	ReplyMessageID                 int           `json:"replyMessageID"`                 // ID of the message to which reply to
	ReplyMessageConversationID     int           `json:"replyMessageConversationID"`     // ID of the message conversation to which reply to
	ForwardedMessageID             int           `json:"forwardedMessageID"`             // ID of the message which is forwarded
	ForwardedMessageConversationID int           `json:"forwardedMessageConversationID"` // ID of the message conversation which is forwarded
}

func SentDateNow() string {
	return time.Now().Format(time.RFC3339)
}
