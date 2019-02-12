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

func ()