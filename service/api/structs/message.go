package structs

import "time"

// A struct which represent a message of a conversation
type Message struct {
	id                 int       //Message id
	body               string    //Message body
	date               time.Time //Message sent time
	picture            string    //Message picture
	sender             int       //User id
	recipient_user     int       //Recipient user id
	recipient_group    int       //Recipient group id
	reply_message      int       //Id of the message to which reply to
	forwareded_message int       //Id of the message which is forwarded
}
