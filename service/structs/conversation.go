package structs

import (
	"regexp"
	"time"
)

type ConversationType int

const (
	PrivateConversationType ConversationType = iota
	GroupConversationType
)

func (recipient_type ConversationType) String() string {
	return [...]string{"PrivateConversationType", "GroupConversationType"}[recipient_type]
}

type ConversationPreview struct {
	ID                  int              `json:"id"`                  // Recipient ID
	Type                ConversationType `json:"type"`                // Recipient type
	RecipientName       string           `json:"recipientName"`       // Recipient name (Group name or user username)
	RecipientPicture    Picture          `json:"recipientPicture"`    // Recipient picture
	LastMessageBody     string           `json:"lastMessageBody"`     // Last message body with the recipient
	LastMessagePicture  string           `json:"lastMessagePicture"`  // Last message picture with the recipient
	LastMessageSentDate time.Time        `json:"lastMessageSentDate"` // Last message sent date with the recipient
}

type Conversation struct {
	ID       int       `json:"id"`       // Conversation ID
	Messages []Message `json:"messages"` // Conversation messages
}

type GroupConversation struct {
	ID            int     `json:"id"`            // Group ID (corresponding to conversation ID)
	Name          string  `json:"name"`          // Group name
	Picture       Picture `json:"picture"`       // Group picture
	UserCreatorID int     `json:"userCreatorID"` // ID of the user who created the group
}

func IsGroupNameValid(groupName string) bool {
	if groupName == "" || (len(groupName) < 3 || len(groupName) > 20) {
		return false
	}
	if !regexp.MustCompile(`^[a-zA-Z0-9\s!?-]{3,20}$`).MatchString(groupName) {
		return false
	}
	return true
}
