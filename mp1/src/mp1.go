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
	//	initHostInformation(mode_remote)
	initHostInformation(mode_local)

	thisID := getHostIndexByPort(listenPort)
	Hosts[thisID].UserName = userName

	hosts_status[thisID] = true

	// Dail to all servers
	go sendServers(listenPort, i_totaluser)

	listenhost := ":" + listenPort

	l, err := net.Listen("tcp", listenhost)
	fmt.Println("listen port now is ", listenPort)
	if err != nil {
		fmt.Println("Listen failed")
		return
	}

	fmt.Println("Listening on localhost : " + listenPort)

	for {

		conn, err := l.Accept()
		if err != nil {
			fmt.Println("accept failed")
			continue
		}
		
		// for remote test
		// // hostId := findHostIndexByConn(conn)
		// // could delete all parameters of port

		// for local test
		hostID := getHostIndexByPort(listenPort)

		Hosts[hostID].Conn = conn
		fmt.Println("after accept=======================", hosts_status, Hosts)
		//hosts_status[hostID] = true
		//fmt.Println(conn.RemoteAddr())
	}
}

func sendServers(port string, n int) {
	// for remote ip test
	// // ipself = getLocalIP()
	count := 1

	for {

		for idx := range Hosts {
			//	fmt.Println("origin count", count)

			// if Hosts[idx].IP_addr == ipself{
			// 	continue
			// }

			if Hosts[idx].Port == port {
				continue
			}
			if hosts_status[idx] == true {
				continue
				// if Hosts[idx].Conn == nil {
				// 	count = count - 1
				// 	hosts_status[idx] = false
				// }
			}
			// for remote ip address
			// //dialAddr := Hosts[idx].IP_addr + ":" + Hosts[idx].Port

			dialAddr := "127.0.0.1:" + Hosts[idx].Port
			dialCon, err := net.Dial("tcp", dialAddr)
			if err == nil {
				fmt.Println("successful connection:-------------------", Hosts[idx].Port)
				count = count + 1
				hosts_status[idx] = true
				Hosts[idx].Conn = dialCon
				fmt.Println(hosts_status, Hosts)
				fmt.Println("connection read,write check=======================", idx, hosts_status, Hosts)
				go readHandler(dialCon, port)

				// for remote version, parameters could be
				// go writeHandler(dialCon)
				go writeHandler(port)
			}
			//fmt.Println("after count", count)
		}
		if count == n {
			break
		}
		//	fmt.Println("exit count", count)
	}
	fmt.Println("READY")
}

func readHandler(conn net.Conn, listenPort string) {

	for {
		var buf = make([]byte, 1024)
		len, err := conn.Read(buf)

		if err != nil {

			// // for remote ip address
			// hostId := getHostIndexByConn(conn)

			hostId := getHostIndexByPort(listenPort)
			left_User := Hosts[hostId].UserName
			fmt.Println(left_User + " has left")
			Hosts[hostId].Conn = nil
			hosts_status[hostId] = false
			fmt.Println("read failture check=======================", hosts_status, Hosts)
			break
		}
		fmt.Printf("Message got from %s is %s\n", listenPort, string(buf[:len]))
	}
}
func writeHandler(listenPort string) {
	// // for remote server
	//index = findHostIndexByConn(conn net.Conn)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		data := scanner.Text()
		for idx := range Hosts {

			if Hosts[idx].Conn != nil && hosts_status[idx] == true {
				fmt.Println("write check=======================", idx, hosts_status, Hosts)
				Hosts[idx].Conn.Write([]byte(data))
			}
		}
		// fmt.Printf("Message got from %s is %s\n", Host[index].UserName, scanner.Text())
		fmt.Printf("Message got from %s is %s\n", listenPort, scanner.Text())
	}
}
