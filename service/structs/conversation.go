package structs

import (
	"regexp"
	"time"
)

type ConversationType int

const (
	PrivateConversation ConversationType = iota
	GroupConversation
)

func (recipient_type ConversationType) String() string {
	return [...]string{"PrivateConversation", "GroupConversation"}[recipient_type]
}

// A struct which represents a conversation
type ConversationPreview struct {
	ID                  int              `json:"id"`                  //Recipient ID
	Type                ConversationType `json:"type"`                //Recipient type
	RecipientName       string           `json:"recipientName"`       //Recipient name (Group name or user username)
	RecipientPicture    string           `json:"recipientPicture"`    //Recipient picture
	LastMessageBody     string           `json:"lastMessageBody"`     //Last message body with the recipient
	LastMessagePicture  string           `json:"lastMessagePicture"`  //Last message picture with the recipient
	LastMessageSentDate time.Time        `json:"lastMessageSentDate"` //Last message sent date with the recipient
}

// A struct which represents a group
type Group struct {
	ID            int    `json:"id"`            //Group ID
	Name          string `json:"name"`          //Group name
	Picture       string `json:"picture"`       //Group picture
	UserCreatorID int    `json:"userCreatorID"` //ID of the user who created the group
}

// Check if a group name is valid
func IsGroupNameValid(groupName string) bool {
	if groupName == "" || (len(groupName) < 3 || len(groupName) > 20) {
		return false
	}
	if !regexp.MustCompile(`^[a-zA-Z0-9\s!?-]{3,20}$`).MatchString(groupName) {
		return false
	}
	return true
}
