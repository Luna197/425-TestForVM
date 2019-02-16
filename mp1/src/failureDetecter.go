package main

import (
	"time"
	"time.Time"
	"time.Duration"
	"sync"
)

/*	failureDetecter
	would give a funciton to server to update and check the value
	use a channel to tell the message of the other hosts 
*/
type failureDetecter struct{
	// detect 
	status *[10]Status
	timeOut [10]Time
	
	// channel
	ch chan string
	// heart beat protocol
	heartBeat_interval time.Duration
}

func (fd *failureDetecter) init( st *[10]Status,hIntval Duration, ch chan string){
	fd.heartBeat_interval = hIntval
	fd.status = st
}

func heartBeat(conn net.Conn, msg chan Message, timeout int) {
     select {
    	 case mg := <- msg:
	// finially receive timestap condition
	  		if mg.timestamp <= timeOut[mg.src]{
				timeOut[msg.src] = msg.timestamp.Add(time.Duration(timeout) * time.Second))
			}
			else{
				//dead
				status[msg.src] = 0
			}
			break
        case <- time.After(10 * time.Second):
	//        Log("conn dead now")
            conn.Close()
}



