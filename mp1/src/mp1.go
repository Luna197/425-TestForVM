package main

import (
	"fmt"
	"os"
	"net"
)

var hosts_status [10]Status

func main(){
	if len(os.Args)< 4 {
		fmt.Println("invalid arguments. please use the following format: ")
		fmt.Println("\"./mp1 name port n\"")
		os.Exit(1)
	}
	//fmt.Printf("arguments : %v\n", os.Args)
	userName, listenPort, totaluser := os.Args[1], os.Args[2], os.Args[3]
	fmt.Printf("Username : %v, Port : %v ,totaluser : %v\n", userName, listenPort, totaluser)

	/* init Process
	  create different channels
		mcast_app_ch : mcast->app layer
			deliver messages to the application layer
		app_mcast_ch : app->mcast layer
			send messages from app to other hosts
	*/
	tcp_mcast_ch := make( chan Message )
	mcast_app_ch := make( chan string )
	app_mcast_ch := make( chan string )
	tcp_fdetect_ch := make( chan Message)

	defer close(tcp_mcast_ch)
	defer close(mcast_app_ch)
	defer close(app_mcast_ch)
	defer close(tcp_fdetect_ch)

	var fdet failureDetecter{}
	fdet.init(&hosts_status,tcp_fdetect_ch)

	mcast multicast := &causal_Multicast{}
	multicast.init( tcp_mcast_ch, mcast_app_ch, app_mcast_ch)

	var client *appLayer
	client.init( mcast_app_ch, app_mcast_ch)
	

	// Listen for incoming connections.
    l, err := net.Listen("tcp", "localhost:" + listenPort )
    exitOnErr(err, "Error listening:", err.Error())

    // Close the listener when the application closes.
    defer l.Close()
    fmt.Println("Listening on localhost : " + listenPort )
    for {
        conn, err := l.Accept()
        exitOnErr(err, "Error accepting: ", err.Error())
        go handleRequest(conn, tcp_mcast_ch)
    }
}

func handleRequest( conn net.Conn, tcp_mcast_ch chan Message ){
	// Make a buffer to hold incoming data.
	buf := make([]byte, 1024)
	// Read the incoming connection into the buffer.
	len, err := conn.Read(buf)
	exitOnErr(err, "Error reading:"err.Error())

	var jsonMsg Message
	err = json.Unmarshal(buf[:len],&jsonMsg)
	exitOnErr(err, "Error Unmarshal data:"err.Error())

	// go multicastMsg(jsonMsg, tcp_mcast_ch)
	// go heartBeat(conn, tcp_mcast_ch, 6)
	// message router
	fmt.Println(jsonMsg)

	switch jsonMsg.type{
		case msg_heartbeat:
			tcp_fdetect_ch <-jsonMsg
			fmt.Printf("recieved Heartbeat: %v", jsonMsg)
		case msg_userMsg:
			fmt.Printf("received User Msg : %v", jsonMsg)
			tcp_mcast_ch <- jsonMsg 
			
		default:
			fmt.Printf("unknownw msg :%v", jsonMsg)
	}
	fmt.Printf("received message(%v): %v \n", len, string(buf[:19]) )
	

	// Send a response back to person contacting us.
	conn.Write([]byte("Message received."))
	// Close the connection when you're done with it.
	conn.Close()
}