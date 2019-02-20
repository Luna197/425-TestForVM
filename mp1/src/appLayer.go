package main

import (
	"fmt"
	"bufio"
	"os"
)
/* appLayer
	Handle user level messages

	What to expect:
		snd_ch: chan to the multicast layer.
			the lower layer would handle all the io
		rcv_ch: chan from the multicast layer.
			messages include the user name .. etc
*/
type appLayer struct {
	// snd_ch <-chan string
	// rcv_ch chan<- string
	rcv_ch <-chan string
	snd_ch chan<- string
}

func (ap *appLayer) init(rch <-chan string, sch chan<- string){
	// set up variales
	ap.rcv_ch = rch
	ap.snd_ch = sch

	// start hendlers
	go ap.handleUserInput()
	go ap.handleMcastInput()
}

func (ap *appLayer) handleMcastInput(){
	for str:= range ap.rcv_ch{
		fmt.Println(str)
	}
}

func (ap *appLayer) handleUserInput(){
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
    	ap.snd_ch <- scanner.Text()
    }
}