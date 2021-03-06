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
	// handle input
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

	/*
		init all Host infor
	*/
	initHostInformation(mode_remote)
	setMyHostInformation(listenPort, userName)

	//fmt.Println(Hosts)
	//initHostInformation(mode_local)
	ipself := getLocalIP()
	thisID := getHostIndexByIP(ipself)
	Hosts[thisID].UserName = userName

	hosts_status[thisID] = true

	// Dail to all servers
	go sendServers(i_totaluser)

	listenHost := ":" + listenPort

	tcpAddr, err := net.ResolveTCPAddr("tcp4", listenHost)
	if err != nil {
		fmt.Println("bug bug bug")
		return
	}

	l, err := net.ListenTCP("tcp", tcpAddr)
	fmt.Println("listen port now is ", listenHost)
	if err != nil {
		fmt.Println("Listen failed")
		return
	}

	for {
		conn, err := l.AcceptTCP()
		if err != nil {
			fmt.Println("accept failed")
			continue
		}
		Hosts[MyHostIndex].Conn = conn
		hosts_status[MyHostIndex] = true
		// 		fmt.Println("after accept=======================", hosts_status[0], hosts_status[1])
		// 		fmt.Println("after accept=======================", Hosts[0], Hosts[1])
		//fmt.Println(conn.RemoteAddr())
	}
}

func sendServers(n int) {
	//for remote ip testg
	//ipself := getLocalIP()
	count := 1

	for {

		for idx := range Hosts {
			if idx == MyHostIndex {
				continue
			}
			//	fmt.Println("origin count", count)

			if hosts_status[idx] == true {
				if Hosts[idx].Conn == nil {
					count = count - 1
					hosts_status[idx] = false
				}
				continue
			}

			//for remote ip address
			dialAddr := Hosts[idx].IP_addr + ":" + Hosts[idx].Port
			//fmt.Println(dialAddr)

			// dialAddr := "127.0.0.1:" + Hosts[idx].Port
			dialCon, err := net.Dial("tcp", dialAddr)
			if err == nil {
				fmt.Println("successful connection:-------------------", Hosts[idx].UserName)
				count = count + 1
				hosts_status[idx] = true
				Hosts[idx].Conn = dialCon

				// 				//fmt.Println(hosts_status, Hosts)
				// 				fmt.Println("after connection and before read,write check=======================", idx)
				// 				fmt.Println("after connection and before read,write check=======================", hosts_status[0], hosts_status[1])
				// 				fmt.Println("after connection and before read,write check=======================",  Hosts[0], Hosts[1])
				go readHandler(dialCon)

				// for remote version, parameters could be
				// go writeHandler(dialCon)
				go writeHandler(idx)
			}
			//fmt.Println("after count", count)
		}
		if count == n {
			
			break
		}
			
	}
	fmt.Println("READY")
}

func readHandler(conn net.Conn) {

	hostId := findHostIndexByConn(conn)
	for {
		var buf = make([]byte, 1024)
		len, err := conn.Read(buf)

		if err != nil {

			// // for remote ip address
			//	fail_hostId := findHostIndexByConn(conn)

			// hostId := getHostIndexByPort(listenPort)
			left_User := Hosts[hostId].UserName
			fmt.Println(left_User + " has left")
			Hosts[hostId].Conn = nil
			hosts_status[hostId] = false
			conn.Close()
			// 			fmt.Println("read failture check=======================",hostId)
			// 			fmt.Println("read failture check=======================",hosts_status[0], hosts_status[1])
			// 			fmt.Println("read failture check=======================",Hosts[0], Hosts[1])
			
			// Continue to Dail to all servers
			break
		}

		fmt.Printf("Message got from %s is %s\n", Hosts[hostId].UserName, string(buf[:len]))
	}
}

func writeHandler(index int) {
	// // for remote server
	//index = findHostIndexByConn(conn net.Conn)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		data := scanner.Text()
		for idx := range Hosts {

			fmt.Println("before write to check the conn and status")
			if Hosts[idx].Conn != nil && hosts_status[idx] == true {
				// 				fmt.Println("write check=======================", idx)
				// 				fmt.Println("write check=======================", hosts_status[0], hosts_status[1])
				// 				fmt.Println("write check=======================", Hosts[0], Hosts[1])
				Hosts[idx].Conn.Write([]byte(data))
			}
		}
		// fmt.Printf("Message got from %s is %s\n", Host[index].UserName, scanner.Text())
		fmt.Printf("Message got from %s is %s\n", Hosts[index].UserName, scanner.Text())
	}
}
