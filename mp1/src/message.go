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
	senderName		string `json:"senderName"` // sender
	senderIdx		int `json:"senderIdx"`

	// Heart beat msg
	timestamp	string `json:"timestamp,omitempty"`

	// Lamport timestamp
	local_timestamp lTimeStamp_t string `json:"local_timestamp, omitempty"`

	// userMsg
	text		string `json:"text,omitempty"`
}

// message heap
type Message_heap struct{
	sender_Idx int
	msg_pts []*Message
}
// implement heap interface
func (msh *Message_heap) Len() int{ return len(msh.msg_pts)}
func (msh *Message_heap) Less(i,j int) bool{
	idx := msh.sender_Idx
	arr := msh.msg_pts 
	return (*arr[i])[idx] < (*arr[j])[idx]
}
func (msh *Message_heap) Swap(i,j int){
	idx := msh.sender_Idx
	arr := msh.msg_pts
	arr[i], arr[j] = arr[j], arr[i]
}
func (msh *Message_heap) Push( x interface{}){ *msh = append(*msh, x.(*Message))}

func (msh *Message_heap) Pop() interface{} {
	old := ts.msg_pts
	n := len(old)
	x := old[n-1] // why
	*ts = old[0 : n-1]
	return x
}

func (msh *Message_heap) getFirstTimeStamp() interface{}{
	return (*msh)msg_pts[0]
}

func (msh *Message_heap) getFirstMessage() interface{}{
	return (*msh)msg_pts
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