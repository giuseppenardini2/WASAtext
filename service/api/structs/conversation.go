package structs

type ConversationType int

const (
	UserConversation ConversationType = iota
	GroupConversation
)

func (conversation_type ConversationType) String() string {
	return [...]string{"UserConversation", "GroupConversation"}[conversation_type]
}

// A struct which represents a conversation
type Conversation struct {
	RecipientID int              `json:"id"`          //Recipient id
	Type        ConversationType `json:"type"`        //Recipient type
	Name        string           `json:"name"`        //Recipient name (Group name or user username)
	Picture     string           `json:"picture"`     //Recipient picture
	LastMessage string           `json:"lastMessage"` //Last message with the recipient
}
