package main

import (
	"fmt"
)

/*
 Implement all of the functionalities of multicast
	including Integrity, Validatiy, Agreement properties
	use two channels to communiate with lower and upper layers.

 // recive channel would get messages from server (lower level)
 // deliver channel would send string to user (upper level)

 Notice: call init funciton before any operation
*/
type multicast interface{
	init()
	// getReceiveChan() ( <-chan Message )
	// getDeliverChan() ( chan<- string )

	// getSendChan() ( chan<- string )
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
		// hold backqueue
		// records of sequence number from different process

}

func (cm *causal_Multicast) init( rch chan Message, dch chan string, sch chan string) {
	cm.rcv_ch = rch
	cm.del_ch = dch
	cm.snd_ch = sch

	// start a new go routine to handle send messages

}

// func (cm *causal_Multicast) getReceiveChan() <-cha n{
// 	return cm.rcv_ch
// }

// func (cm *causal_Multicast) getDeliverChan() chan<- {
// 	return cm.del_ch
// }