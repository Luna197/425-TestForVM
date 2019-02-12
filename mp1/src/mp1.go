package main

import (
	"fmt"
	"os"
	"net"
)

func main(){
	if len(os.Args)< 4 {
		fmt.Println("invalid arguments. please use the following format: ")
		fmt.Println("\"./mp1 name port n\"")
		os.Exit(1)
	}
	//fmt.Printf("arguments : %v\n", os.Args)
	userName, listenPort, totaluser := os.Args[1], os.Args[2], os.Args[3]
	fmt.Printf("Username : %v, Port : %v ,totaluser : %v\n", userName, listenPort, totaluser)

	// init Process


	// Listen for incoming connections.
    l, err := net.Listen("tcp", "localhost:" + listenPort )
    exitOnErr(err, "Error listening:", err.Error())

    // Close the listener when the application closes.
    defer l.Close()
    fmt.Println("Listening on localhost : " + listenPort )
    for {
        conn, err := l.Accept()
        exitOnErr(err, "Error accepting: ", err.Error())
        go handleRequest(conn)
    }
}

func handleRequest( conn net.Conn){
	// Make a buffer to hold incoming data.
	buf := make([]byte, 1024)
	// Read the incoming connection into the buffer.
	len, err := conn.Read(buf)
	exitOnErr(err, "Error reading:"err.Error())

	var jsonMsg Message
	err = json.Unmarshal(buf[:len],&jsonMsg)
	exitOnErr(err, "Error Unmarshal data:"err.Error())

	// handle Message based on type
	fmt.Println(jsonMsg)
	switch jsonMsg.type{
		case msg_heartbeat:
			fmt.Printf("recieved Heartbeat: %v", jsonMsg)
		case msg_userMsg:
			fmt.Printf("received User Msg : %v", jsonMsg)
		default:
			fmt.Printf("unknownw msg :%v", jsonMsg)
	}
	//fmt.Printf("received message(%v): %v \n", len, string(buf[:19]) )
	

	// Send a response back to person contacting us.
	conn.Write([]byte("Message received."))
	// Close the connection when you're done with it.
	conn.Close()
}