// THis file is for testing
package main

import (
	"fmt"
	"net"
)

func main() {

	// test local IP
	printTestString("test local ip", 0)
	fmt.Printf("local IP: %v \n", getLocalIP())

	// test json file load successfully
// <<<<<<< HEAD
// 	printTestString("test json file", 0)
// 	hs := getRemoteServers()
// 	fmt.Printf("Total number of servers: %2v (should be 9 on the server)\n", len(hs))
// =======
	printTestString("test json file",0)
	initHostInformation( mode_local)
	fmt.Printf("Total number of servers: %2v (should be 9 on the server)\n", len(Hosts))
	for _,h := range Hosts{
		fmt.Println(h)
	}


	testAddr := "localhost"
	idx := getHostIndexByIP(testAddr)
	fmt.Printf("lookup ip: %v -> its host :%v\n", testAddr, idx)


	conn, err := net.Dial("udp", "8.8.8.8:80")
	idx = findHostIndexByConn(conn)
	exitOnErr(err,"cannot connect to 8.8.8.8:80:")
	defer conn.Close()
	fmt.Printf("lookup conn: %v -> its host :%v\n", conn, idx)
// >>>>>>> d7006056134c6d97536a6a1513939b3685343f06

	for h := range hs{
		fmt.Printf()
	}

	fmt.Println("Test complete")
}

func printTestString(str string, indent int) {
	for i := 0; i < indent; {
		fmt.Printf("\t")
	}
	fmt.Printf("===== %v =====\n", str)
}
