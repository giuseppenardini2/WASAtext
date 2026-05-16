package structs

import "time"

// A struct which represents a conversation
type Conversation struct {
	RecipientID         int           `json:"recipientId"`         //Recipient ID
	RecipientType       RecipientType `json:"recipientType"`       //Recipient type
	RecipientName       string        `json:"recipientName"`       //Recipient name (Group name or user username)
	RecipientPicture    string        `json:"recipientPicture"`    //Recipient picture
	LastMessageBody     string        `json:"lastMessageBody"`     //Last message body with the recipient
	LastMessagePicture  string        `json:"lastMessagePicture"`  //Last message picture with the recipient
	LastMessageSentDate time.Time     `json:"lastMessageSentDate"` //Last message sent date with the recipient
}
