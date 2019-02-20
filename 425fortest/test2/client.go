package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
	"bufio"
)
// 127.0.0.1:444444
// 127.0.0.1:555555

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		// read from the connection
		var buf = make([]byte, 10)
		// log.Println("start to read from conn")

		n, err := c.Read(buf)
		if err != nil {
			log.Println("failture of this connection", err)
			return
		}
	    log.Printf(" %s\n", string(buf[:n]))
		// data := string(buf[:n])

				
		// for  key, val := range connections{
		// 	// log.Printf("wtf", key,  "111")
		// 	data += key
		// 	val.Write([]byte(data))
		//    	//  val.Write(string(data))
		// }

	}
}

func main() {
    // if len(os.Args) <= 1 {
    //     fmt.Println("usage: go run client2.go YOUR_CONTENT")
    //     return
    // }
    log.Println("begin dial...")
    conn, err := net.Dial("tcp", ":8888")
    go handleConn(conn)
    
    log.Println("focus!", conn.RemoteAddr().String())
    if err != nil {
        log.Println("dial error:", err)
        return
    }
    defer conn.Close()
    log.Println("dial ok")

    time.Sleep(time.Second * 2)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		data := scanner.Text()
  	    conn.Write([]byte(data))
	    fmt.Println(scanner.Text())
	}

    // data := os.Args[1]
    // conn.Write([]byte(data))

    time.Sleep(time.Second * 10000)
}