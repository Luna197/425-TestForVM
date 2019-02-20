package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
)

var hosts_status [10]bool

func main() {
	if len(os.Args) < 4 {
		fmt.Println("invalid arguments. please use the following format: ")
		fmt.Println("\"./mp1 name port n\"")
		os.Exit(1)
	}
	// //fmt.Printf("arguments : %v\n", os.Args)
	userName, listenPort, totaluser := os.Args[1], os.Args[2], os.Args[3]
	fmt.Printf("Username : %v, Port : %v ,totaluser : %v\n", userName, listenPort, totaluser)

	i_totaluser, err := strconv.Atoi(totaluser)
	if err != nil {
		exitOnErr(err, "string to int conver fail")
	}

	initHostInformation(mode_local)

	thisID := getHostIndexByPort(listenPort)
	Hosts[thisID].UserName = userName

	hosts_status[thisID] = true

	go sendServers(listenPort, i_totaluser)
	/* init Process
	  create different channels
		mcast_app_ch : mcast->app layer
			deliver messages to the application layer
		app_mcast_ch : app->mcast layer
			send messages from app to other hosts
	*/
	// tcp_mcast_ch := make( chan Message )
	// mcast_app_ch := make( chan string )
	// app_mcast_ch := make( chan string )
	// //tcp_fdetect_ch := make( chan Message)

	// defer close(tcp_mcast_ch)
	// defer close(mcast_app_ch)
	// defer close(app_mcast_ch)
	//defer close(tcp_fdetect_ch)

	// 	var fdet failureDetecter
	// 	fdet.init(&hosts_status,tcp_fdetect_ch)
	// >>>>>>> master

	// mcast multicast := &causal_Multicast
	// multicast.init( tcp_mcast_ch, mcast_app_ch, app_mcast_ch)

	// var client *appLayer
	// client.init( mcast_app_ch, app_mcast_ch)

	// for h := range Hosts {

	// 	fmt.Println(Hosts[h].Port)
	// }

	//Listen for incoming connections.
	//"localhost:"
	listenhost := ":" + listenPort

	l, err := net.Listen("tcp", listenhost)
	fmt.Println("listen port now is ", listenPort)
	if err != nil {
		fmt.Println("Listen failed")
		return
	}

	//exitOnErr(err, "Error listening:" + err.Error())

	// Close the listener when the application closes.
	//defer l.Close()

	fmt.Println("Listening on localhost : " + listenPort)

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("accept failed")
			continue
		}
		//go Handler(conn, Hosts)

		hostID := getHostIndexByPort(listenPort)
		// if hostID == -1 {
		// 	fmt.Println("Cnnot find host index")
		// 	os.Exit(-1)
		// }
		//	hosts_status[hostID] = true
		Hosts[hostID].Conn = conn
	}
}

func sendServers(port string, n int) {
	count := 1
	for {
		//88fmt.Println(count)
		// loop all ips  in hosts
		for idx := range Hosts {
			//	fmt.Println("----------------------------",idx, port, Hosts[idx].Port)
			if Hosts[idx].Port == port {
				continue
			}
			if hosts_status[idx] == true {
				if Hosts[idx].Conn == nil {
					count = count - 1
					hosts_status[idx] = false
				}
			}
			dialAddr := "127.0.0.1:" + Hosts[idx].Port
			dialCon, err := net.Dial("tcp", dialAddr)
			if err == nil {
				fmt.Println("successful connection:-------------------", Hosts[idx].Port)
				count = count + 1
				hosts_status[idx] = true
				Hosts[idx].Conn = dialCon
				go readHandler(dialCon, port)
				go writeHandler(port)
			}
		}
		if count == n {
			break
		}
	}
	fmt.Println("READY")
}

func readHandler(conn net.Conn, listenPort string) {

	for {
		var buf = make([]byte, 1024)
		//fmt.Println("handle connection ============= ")
		len, err := conn.Read(buf)

		if err != nil {
			hostId := getHostIndexByPort(listenPort)

			// should get user name from the connection and update
			left_User := Hosts[hostId].UserName
			fmt.Println(left_User + " has left")
			Hosts[hostId].Conn = nil
			hosts_status[hostId] = false
		//	conn.Close()
			break
		}
		fmt.Printf("Message got from %s is %s\n", listenPort, string(buf[:len]))
	}
}
func writeHandler(listenPort string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		data := scanner.Text()
		for idx := range Hosts {

			if Hosts[idx].Conn != nil {
				Hosts[idx].Conn.Write([]byte(data))
			}
		}
		fmt.Printf("Message got from %s is %s\n", listenPort, scanner.Text())
	}
}

// // // for - accept,
// // func handleRequest( conn net.Conn, tcp_mcast_ch chan Message ){
// // 	// Make a buffer to hold incoming data.
// // 	buf := make([]byte, 1024)
// // 	// Read the incoming connection into the buffer.
// // 	len, err := conn.Read(buf)
// // 	exitOnErr(err, "Error reading:" + err.Error())
// // 	//find dead and delet from muliticast
// // //if conn failed in here, just regard this user left
// // // conn faild - conn, if, user - failed - left

// // 	var jsonMsg Message
// // 	err = json.Unmarshal(buf[:len],&jsonMsg)
// // 	exitOnErr(err, "Error Unmarshal data:" + err.Error())

// // 	// go multicastMsg(jsonMsg, tcp_mcast_ch)
// // 	// go heartBeat(conn, tcp_mcast_ch, 6)
// // 	// message router
// // 	fmt.Println(jsonMsg)

// // 	switch jsonMsg.msg_type{
// // 		case msg_heartbeat:
// // 			// tcp_fdetect_ch <-jsonMsg
// // 			// fmt.Printf("recieved Heartbeat: %v", jsonMsg)
// // 		case msg_userMsg:
// // 			fmt.Printf("received User Msg : %v", jsonMsg)
// // 			tcp_mcast_ch <- jsonMsg

// // 		default:
// // 			fmt.Printf("unknownw msg :%v", jsonMsg)
// // 	}
// // 	fmt.Printf("received message(%v): %v \n", len, string(buf[:19]) )

// // 	// Send a response back to person contacting us.
// // 	conn.Write([]byte("Message received."))
// // 	// Close the connection when you're done with it.
// // 	conn.Close()
