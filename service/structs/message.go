package structs

import (
	"errors"
	"time"
)

type RecipientType int

const (
	RecipientUser RecipientType = iota
	RecipientGroup
)

func (recipient_type RecipientType) String() string {
	return [...]string{"RecipientUser", "RecipientGroup"}[recipient_type]
}

// A struct which represent a message of a conversation
type Message struct {
	ID                 int       `json:"id"`                 //Message ID
	Body               string    `json:"body"`               //Message body
	SentDate           time.Time `json:"sent date"`          //Message sent time
	Picture            string    `json:"picture"`            //Message picture
	SenderID           int       `json:"senderID"`           //Sender ID
	RecipientUserID    int       `json:"recipientUserID"`    //Recipient user ID
	RecipientGroupID   int       `json:"recipientGroupID"`   //Recipient group ID
	ReplyMessageID     int       `json:"replyMessageID"`     //ID of the message to which reply to
	ForwardedMessageID int       `json:"forwardedMessageID"` //ID of the message which is forwarded
}

func (message Message) GetRecipientType() (RecipientType, error) {
	if message.RecipientUserID != -1 {
		return RecipientUser, nil
	}
	if message.RecipientGroupID != -1 {
		return RecipientGroup, nil
	}
	return RecipientGroup, errors.New("Recipient type error")
}
