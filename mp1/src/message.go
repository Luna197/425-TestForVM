package main

import (
	"fmt"
	"time"
)

type MessageType_t int

const (
	msg_none      MessageType_t = 0
	msg_heartbeat MessageType_t = 1
	msg_userMsg   MessageType_t = 2
)

// Sned all the data throught json
<<<<<<< HEAD
type Message struct {
	msg_type   MessageType_t
	senderName string `json:"senderName"` // sender
	senderIdx  int    `json:"senderIdx"`

	// Heart beat msg
	timestamp time.Time `json:"timestamp,omitempty"`

	// Lamport timestamp
	local_timestamp lTimeStamp_t `json:"local_timestamp,omitempty"`

	// userMsg
	text string `json:"text,omitempty"`
}

// message heap
type Message_heap struct {
	sender_Idx int
	msg_pts    []*Message
=======
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
>>>>>>> 3deab4c06c30b3e08ada80f01728d6570f30c802
}

// implement heap interface
<<<<<<< HEAD
func (msh *Message_heap) Len() int { return len(msh.msg_pts) }
func (msh *Message_heap) Less(i, j int) bool {
	idx := msh.sender_Idx
	arr := msh.msg_pts
	return (*arr[i])[idx] < (*arr[j])[idx]
}
func (msh *Message_heap) Swap(i, j int) {
	idx := msh.sender_Idx
	arr := msh.msg_pts
	arr[i], arr[j] = arr[j], arr[i]
}
func (msh *Message_heap) Push(x interface{}) { *msh = append(*msh, x.(*Message)) }
=======
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
>>>>>>> 3deab4c06c30b3e08ada80f01728d6570f30c802

func (msh *Message_heap) Pop() interface{} {
	old := msh.msg_pts
	n := len(msh.msg_pts)
	x := msh.msg_pts[n-1]
	msh.msg_pts = old[0 : n-1]
	return x
}

<<<<<<< HEAD
func (msh *Message_heap) getFirstTimeStamp() interface{} {
	return *msh.msg_pts[0]
}

func (msh *Message_heap) getFirstMessage() interface{} {
	return *msh.msg_pts
=======
func (msh *Message_heap) getFirstTimeStamp() interface{}{
	return *msh.msg_pts[0]
}

func (msh *Message_heap) getFirstMessage() interface{}{
	return msh.msg_pts[0]
>>>>>>> 3deab4c06c30b3e08ada80f01728d6570f30c802
}

type MsgHandler interface {
	handle(message string)
}

func (msg Message) String() string {
	var typeStr string
<<<<<<< HEAD
	switch msg.msg_type {
	case msg_heartbeat:
		typeStr = "Hrt"
		return fmt.Srintf("<Message type=%v , timestamp:%v>", typeStr, h.timestamp)
	case msg_userMsg:
		typeStr = "Msg"
		var shortText string
		if len(h.text) > 10 {
			shortText = h.text[:10] + "..."
		} else {
			shortText = h.text
		}
		return fmt.Srintf("<Message type=%v , text:%v>", typeStr, shortText)
	default:
		return fmt.Srintf("Message : unknown type")
=======
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
>>>>>>> 3deab4c06c30b3e08ada80f01728d6570f30c802
	}
}

// Useful functions
<<<<<<< HEAD
//func
=======
//func 
>>>>>>> 3deab4c06c30b3e08ada80f01728d6570f30c802
