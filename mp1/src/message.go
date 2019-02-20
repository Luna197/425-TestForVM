package main

import (
	"fmt"
)

type MessageType_t int
const(
	msg_none		MessageType_t = 0
	msg_heartbeat	MessageType_t = 1
	msg_userMsg		MessageType_t = 2 
)

// Sned all the data throught json
type Message struct{
	MsgType	MessageType_t
	SenderName		string `json:"SenderName"` // sender
	SenderIdx		int `json:"SenderIdx"`

	// Lamport timestamp
	LocalTimeStamp lTimeStamp_t `json:"LocalTimeStamp, omitempty"`

	// userMsg
	Text		string `json:"Text,omitempty"`
}

// message heap
type Message_heap struct{
	SenderIdx int
	msg_pts []*Message
}
// implement heap interface
func (msh Message_heap) Len() int{ return len(msh.msg_pts)}
func (msh Message_heap) Less(i,j int) bool{
	idx := msh.SenderIdx
	arr := msh.msg_pts 
	return arr[i].LocalTimeStamp[idx] < arr[j].LocalTimeStamp[idx]
}
func (msh Message_heap) Swap(i,j int){
	arr := msh.msg_pts
	arr[i], arr[j] = arr[j], arr[i]
}
func (msh *Message_heap) Push( x interface{}){
	msh.msg_pts = append( msh.msg_pts, x.(*Message))
}

func (msh *Message_heap) Pop() interface{} {
	old := msh.msg_pts
	n := len(msh.msg_pts)
	x := msh.msg_pts[n-1]
	msh.msg_pts = old[0 : n-1]
	return x
}

func (msh *Message_heap) getFirstTimeStamp() interface{}{
	return *msh.msg_pts[0]
}

func (msh *Message_heap) getFirstMessage() interface{}{
	return msh.msg_pts[0]
}



type MsgHandler interface{
	handle( message string )
	
}

func (msg Message)String() string{
	var typeStr string
	switch msg.MsgType{
		case msg_userMsg:
			typeStr = "Msg"
			var shortText string
			if len(msg.Text) > 10{
				shortText = msg.Text[:10]+"..."
			}else{
				shortText = msg.Text
			}
			return fmt.Sprintf("<Message type=%v , Text:%v>", typeStr, shortText)
		default:
			return fmt.Sprintf("Message : unknown type")
	}
}

// Useful functions
//func 