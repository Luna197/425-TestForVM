package main

import (
	"fmt"
	"time"
)

type Status_t int

const (
	status_dead  MessageType_t = 0
	status_alive MessageType_t = 1
)

/*	failureDetecter
	would give a funciton to server to update and check the value
	use a channel to tell the message of the other hosts
*/
type failureDetecter struct {
	// detect
	status  *[10]Status_t
	timeOut [10]Time

	// channel to receive heartbeat message
	recv_ch <-chan string

	// heart beat protocol
	heartBeat_interval time.Duration
}

func (fd *failureDetecter) init(st *[10]Status, hIntval Duration, ch chan string) {
	fd.heartBeat_interval = hIntval
	fd.status = st

	go fd.startHeartBeating()
	go fd.recvMsg_handler()
	go fd.checkTimeOut()
}

func (fd *failureDetecter) startHeartBeating() {
	var msg Message
	msg.msg_type = msg_heartbeat
	for {
		msg.timestamp = time.Now()
		multicastMsg(msg, false)
		// need to check time again
		time.Sleep(100 * time.Millisecond)
	}
}

func (fd *failureDetecter) checkTimeOut() {
	var last_status [10]Status_t
	var curTime time.Time

	for {
		curTime = time.Now()
		for idx, deadline := range fd.timeOut {
			if curTime >= deadline {
				if last_status[idx] == status_alive {
					fmt.Printf("user %v leaved\n", idx)
				}
				fd.status[idx] = status_dead
			} else {
				if last_status[idx] == status_dead {
					fmt.Printf("user %v is joined\n", idx)
				}
				fd.status[idx] = status_alive
			}
		}
		// record last status to local storage
		for idx, st := range fd.status {
			last_status[idx] = st
		}
		time.Sleep(10 * time.Millisecond) // to be change
	}
}

// receive message and check hearBest fail or not, update timeout
func (fd *failureDetecter) recvMsg_handler() {
	for {
		if msg, ok := <-fd.recv_ch; ok {
			if mg.timestamp <= timeOut[mg.src] {
				timeOut[msg.src] = msg.timestamp.Add(time.Duration(timeout) * time.Second)
			}
		} else {
			break
		}

	}
}
