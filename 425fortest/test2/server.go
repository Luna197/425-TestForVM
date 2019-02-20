
package main

import (
	"log"
	"net"
)

func handleConn(c net.Conn, connections map[string]net.Conn) {
	defer c.Close()
	for {
		// read from the connection
		var buf = make([]byte, 10)
		log.Println("start to read from conn")

		n, err := c.Read(buf)
		if err != nil {
			log.Println("failture of this connection", err)
			c.Close()
			delete(connections, c.RemoteAddr().String())
			return
		}
	    log.Printf(" %s\n", string(buf[:n]))
		data := string(buf[:n])

				
		for  key, val := range connections{
			// log.Printf("wtf", key,  "111")
			key = key
			val.Write([]byte(data))
		   	//  val.Write(string(data))
		}
	}
}
//127.0.0.1:58664
//checkAlive[10]


func main() {

	connections := make(map[string]net.Conn)

	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Println("listen error:", err)
		return
	}

	for {
		conn, err := l.Accept()
		ipnow := conn.RemoteAddr().String()
		log.Println("server dection:", conn.RemoteAddr().String())
		if err != nil {
			log.Println("accept error:", err)

			break
		}
		connections[ipnow] = conn
		// start a new goroutine to handle
		// the new connection.
		log.Println("accept a new connection")

		go handleConn(conn, connections)

		
	}
}