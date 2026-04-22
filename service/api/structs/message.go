package structs

import(
	"time"
)

//Message
type Message struct {
	id int
	body string
	date time.Time
	picture string
	sender int
	recipient_user int
	recipient_group int
	reply_message int
	forwareded_message int
}