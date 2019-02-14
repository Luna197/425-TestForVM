package main

import (
	"fmt"
)

type MessageType_t int
const(
	msg_heartbeat	MessageType_t = 0
	msg_userMsg		MessageType_t = 1 
)

// Sned all the data throught json
type Message struct{
	msg_type	MessageType_t
	src		string `json:"src"` // sender

	// Heart beat msg
	timestamp	string `json:"timestamp,omitempty"`

	// userMsg
	text		string `json:"text,omitempty"`
}


type MsgHandler interface{
	handle( message string )
	
}

func (msg Message)String() string{
	var typeStr string
	switch msg.msg_type{
		case msg_heartbeat:
			typeStr = "Hrt"
			return fmt.Srintf("<Message type=%v , timestamp:%v>", typeStr, h.timestamp)
		case msg_userMsg:
			typeStr = "Msg"
			var shortText := string
			if len(h.text) > 10{
				shortText = h.text[:10]+"..."
			}else{
				shortText = h.text
			}
			return fmt.Srintf("<Message type=%v , text:%v>", typeStr, shortText)
		default:
			return fmt.Srintf("Message : unknown type")
	}
}

// Useful functions
func 