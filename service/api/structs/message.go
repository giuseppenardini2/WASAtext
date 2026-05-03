package structs

import "time"

// A struct which represent a message of a conversation
type Message struct {
	ID                 int       `json:"id"`                 //Message id
	Body               string    `json:"body"`               //Message body
	Date               time.Time `json:"date"`               //Message sent time
	Picture            string    `json:"picture"`            //Message picture
	SenderID           int       `json:"senderID"`           //Sender user id
	RecipientUserID    int       `json:"recipientUserID"`    //Recipient user id
	RecipientGroupID   int       `json:"recipientGroupID"`   //Recipient group id
	ReplyMessageID     int       `json:"replyMessageID"`     //Id of the message to which reply to
	ForwardedMessageID int       `json:"forwardedMessageID"` //Id of the message which is forwarded
}
