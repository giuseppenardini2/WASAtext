package structs

import (
	"time"
)

// A struct which represent a message of a conversation
type Message struct {
	ID                 int       `json:"id"`                 //Message ID
	ConversationID     int       `json:"conversationID"`     //Conversation ID
	Body               string    `json:"body"`               //Message body
	SentDate           time.Time `json:"sent date"`          //Message sent time
	Picture            string    `json:"picture"`            //Message picture
	SenderID           int       `json:"senderID"`           //Sender ID
	ReplyMessageID     int       `json:"replyMessageID"`     //ID of the message to which reply to
	ForwardedMessageID int       `json:"forwardedMessageID"` //ID of the message which is forwarded
}
