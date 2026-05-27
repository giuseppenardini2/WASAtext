package structs

var acceptedEmojis = []string{
	"👍", "👍",
	"🙂", "🙁", "😭", "😍", "😎", "😏",
	"❤️", "💔",
	"🙉", "🙊", "🙈",
	"💩", "👽"}

// A struct which represent a message of a conversation
type Reaction struct {
	UserID         int    `json:"userID"`         // User ID
	ConversationID int    `json:"conversationID"` // Conversation ID
	MessageID      int    `json:"messageID"`      // Message ID
	Emoji          string `json:"emoji"`          // Emoji
}

func IsReactionValid(emoji string) bool {
	if len(emoji) != 1 {
		return false
	}
	for i := range acceptedEmojis {
		if emoji == acceptedEmojis[i] {
			return true
		}
	}
	return false
}
