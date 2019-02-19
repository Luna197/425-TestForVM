package main

import (
	"fmt"
	"container/heap"
	"os"
)


type lTimeStamp_t []int64


/*
 Implement all of the functionalities of multicast
	including Integrity, Validatiy, Agreement properties
	use two channels to communiate with lower and upper layers.

 // recive channel would get messages from server (lower level)
 // deliver channel would send string to user (upper level)

 Notice: call init funciton before any operation
*/
type multicast interface {
	init()
}

/*
	implement the multicast with
	Casual ordering

	functionalities to implement
		// on receive
			1. hold back the out-of-order messages
			2. filter redundant messages from other host
			3. broadcast new received messages to other host
				to achieve agreement properties
		// on send
			1. send data throught tcp to all alive hosts
			2. must upadta self's lamport timestamps before send
			3. user would pass a string to the struct, and mcast shoudl pack it into
				Message format and send it to other hosts thorught tcp.
		// on deliver
			1. append other user's Id in the string(easier to implement)
*/
type causal_Multicast struct {
	// the two channels

	// for revice
	rcv_ch <-chan Message
	del_ch chan<- string

	// for send
	snd_ch chan<- string

	// internal datastucture
	local_timestamp lamportTimeStamp
	holdback_queues []lTimeStampInfo_heap
}


func (cm *causal_Multicast) init( numHosts int, rch chan Message, dch chan string, sch chan string) {
	cm.ltimestamp = make( lamportTimeStamp, numHosts)
	cm.rcv_ch = rch
	cm.del_ch = dch
	cm.snd_ch = sch

	// start a new go routine to handle send messages
	go recvMsg_handler()
	go sendMsg_handler()
}

/*
	deliver message to user channel and update the local-timestamp
*/
func (cm *causal_Multicast) deliverMsg( msg *Message){
	deliver_str := msg.senderName + ": " + msg.txt
	del_ch <- deliver_str
	// update timestamp
	for i:= len(cm.local_timestamp); i>=0; i--{
		if cm.local_timestamp[i] < msg.local_timestamp[i]{
			cm.local_timestamp[i] = msg.local_timestamp[i] 
		}
	}
}

func (cm *causal_Multicast ) recvMsg_handler(){
	// get message from lower layer
	for msg := range cm.rcv_ch{
		n = len(cm.local_timestamp)
		cts, mts := cm.local_timestamp, msg.local_timestamp

		//check duplicate message
		msg_duplicate := true
		for i:=0; i<n; i++{
			if mts[i] > cts[i]{

				msg_duplicate = false
				break 
			} 
		}
		if msg_duplicate{
			continue
		}

		// new message, broadcast to everyone
		if msg_future || msg_next {
			// broadcast to everyone
			multicastMsg( msg, true)
		}

		// put message in queue and deliver
		heap.Push(cm.holdback_queues[msg.senderIdx], cm)
		// deliver all possible messages and update timestamp
		for delivered_count:=1 ; delivered_count>0;{
			delivered_count = 0
			for _,que := range cm.holdback_queues{

				// message timestamp & local timestamp
				mts := que.getFirstTimeStamp().(lTimeStamp_t)
				lts := cm.local_timestamp
				
				var msg_next, msg_future bool
				for i, next_cnt:=0,0 ; i<n; i++{
					if lts[i] >= mts[i]{
						continue
					}else if lts[i]+1 == mts[i]{
						msg_next = true
						next_cnt++
						if next_cnt>=2{
							msg_future,msg_next = true, false
							break
						}
					}else{
						msg_future,msg_next = true, false
						break
					}
				}
	
				if msg_future {
					continue
				}else if msg_next{
					//deliver msg
					msg_to_be_delivered := heap.Pop(que).(Message)
					cm.deliverMsg(msg_to_be_delivered)
					delivered_count++
				}
			}
		}  
	}
}

func (cm *causal_Multicast ) sendMsg_handler(){
	// get hosts to send message
	// only send to soemone who is alive
	text <- cm.snd_ch.(string)
	var msg Message
	msg.msg_type = msg_userMsg
	msg.local_timestamp = cm.local_timestamp
	msg.text = text
	multicastMsg( msg, true )
}

/*
	Multicast a Message to another servers
	Only append senderName, senderIdx to the message
	the type must be inserted before calling this fuction
*/
func multicastMsg(msg Message, sendOnlyAlive bool) {
	if msg.msg_type = msg_none{
		fmt.Println("multicast message without given a msg_type")
		os.Exit(1)
	}
	msg.senderName =  "to be added" // some name
	msg.senderIdx = 0// some index

	hosts = getRemoteServers()
	// hosts => msg.src
	for _,h:= range hosts{
		if !sendOnlyAlive | (sendOnlyAlive && hosts_status[msg.src]==status_alive){
			conn, err := net.Dial("udp", h)
			defer conn.Close()
			exitOnErr(err, "message connection failed")
			conn.Write(snd_ch <- msg)
		}			
}


