package structs

import (
	"regexp"
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
	ID                    int              `json:"id"`                    // Conversation ID
	Type                  ConversationType `json:"type"`                  // Conversation type
	RecipientID           int              `json:"recipientID"`           // Recipient ID
	RecipientName         string           `json:"recipientName"`         // Recipient name (Group name or user username)
	RecipientPictureURI   string           `json:"recipientPictureURI"`   // Recipient picture URI
	LastMessageID         int              `json:"lastMessageID"`         // Last message ID
	LastMessageBody       string           `json:"lastMessageBody"`       // Last message body with the recipient
	LastMessagePictureURI string           `json:"lastMessagePictureURI"` // Last message picture with the recipient
	LastMessageSentDate   string           `json:"lastMessageSentDate"`   // Last message sent date with the recipient
}

type Conversation struct {
	ID       int       `json:"id"`       // Conversation ID
	Messages []Message `json:"messages"` // Conversation messages
}

var groupConversationNameRegex = regexp.MustCompile(`^[a-zA-Z0-9\s!?-]{3,20}$`)

type GroupConversation struct {
	ID            int    `json:"id"`            // Group ID (corresponding to conversation ID)
	Name          string `json:"name"`          // Group name
	PictureURI    string `json:"pictureURI"`    // Group picture URI
	UserCreatorID int    `json:"userCreatorID"` // ID of the user who created the group
}

func IsGroupNameValid(groupName string) bool {
	if groupName == "" || (len(groupName) < 3 || len(groupName) > 20) {
		return false
	}
	if !groupConversationNameRegex.MatchString(groupName) {
		return false
	}
	return true
}
