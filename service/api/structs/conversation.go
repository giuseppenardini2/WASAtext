package structs

import "fmt"

type ConversationType int

const (
	UserConversation ConversationType = iota
	GroupConversation
)

func (conversation_type ConversationType) String() string {
	return [...]string{"UserConversation", "GroupConversation"}[conversation_type]
}

//A struct which represents a group
type Conversation struct {
	recipient_id int //Recipient id
	conversation_type ConversationType //Recipient type
	name string //Recipient name (Group name or user username)
	picture string //Recipient picture
	last_message string //Last message with the recipient
}