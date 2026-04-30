package structs

import "time"

// A struct which represent a message of a conversation
type Message struct {
	Id               int       `json:"id"`               //Message id
	Body             string    `json:"body"`             //Message body
	Date             time.Time `json:"date"`             //Message sent time
	Picture          string    `json:"picture"`          //Message picture
	Sender           int       `json:"sender"`           //Sender user id
	RecipientUser    int       `json:"recipientUser"`    //Recipient user id
	RecipientGroup   int       `json:"recipientGroup"`   //Recipient group id
	ReplyMessage     int       `json:"replyMessage"`     //Id of the message to which reply to
	ForwardedMessage int       `json:"forwardedMessage"` //Id of the message which is forwarded
}
