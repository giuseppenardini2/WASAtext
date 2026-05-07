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
	RecipientID         int              `json:"recipientId"`         //Recipient id
	ConversationType    ConversationType `json:"conversationType"`    //Recipient type
	RecipientName       string           `json:"recipientName"`       //Recipient name (Group name or user username)
	RecipientPicture    string           `json:"recipientPicture"`    //Recipient picture
	LastMessageBody     string           `json:"lastMessageBody"`     //Last message body with the recipient
	LastMessagePicture  string           `json:"lastMessagePicture"`  //Last message picture with the recipient
	LastMessageSentDate string           `json:"lastMessageSentDate"` //Last message sent date with the recipient
}
