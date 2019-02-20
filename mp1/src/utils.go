package main

import (
	"fmt"
	"net"
	"os"
	"encoding/json"
	"io/ioutil"
)



type RunMode_t int
const(
	mode_local		RunMode_t = 0
	mode_remote		RunMode_t = 1
)

type Host struct{
	Id			string	`json:"id,omitempty"`
	Domain_name	string	`json:"dname,omitempty"`
	IP_addr		string	`json:"ip,omitempty"`
// 	Port		string	`json:"port,omitempty"`
// <<<<<<< HEAD
// 	//conn      net.Conn 
// =======
	UserName	string
	Conn 		net.Conn

// >>>>>>> d7006056134c6d97536a6a1513939b3685343f06
}

// use `initHostInformation()` to initialize this function
var Hosts []Host

// indicate the mode that current running
// 	do not used in other files
var UTILS_currRunMode RunMode_t

func (h Host) String() string{
	return fmt.Sprintf("<Host id:%v, dame:%v, ip:%v, port:%v, conn:%v>", h.Id, h.Domain_name, h.IP_addr, h.Port, h.Conn)
}

func exitOnErr(err error, str string){
	if err != nil{
		fmt.Printf("%v%v", str, err.Error())
		os.Exit(1)
	}
}

func getLocalIP() string{
	conn, err := net.Dial("udp", "8.8.8.8:80")
	exitOnErr(err,"cannot connect to 8.8.8.8:80:")
    
    defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	host, _, err  := net.SplitHostPort(localAddr.String())
	exitOnErr(err, "Cannot split local IP:")
	return host
}

// <<<<<<< HEAD



// func getRemoteServers() []Host{
	
// 	jData, err := ioutil.ReadFile("servers.json")
// =======
/*
	Search the Host array and return the index of matching host
	Used in accepting a new connection
*/
func findHostIndexByConn( conn net.Conn) int{
	if conn == nil{
		return -1
	}
	remoteAddr := conn.RemoteAddr()
	rhost, rport,_  := net.SplitHostPort(remoteAddr.String())
	if UTILS_currRunMode == mode_local{
		for i:=0 ;i<len(Hosts); i++{
			fmt.Printf("checl port:%v, cur port: %v\n", Hosts[i].Port, rport)
			if( rport == Hosts[i].Port ){
				return i
			}
		}
		return -1
	}else{
		return getHostIndexByIP(rhost)
	}
	//fmt.Printf("%v, %v", host, port)

}

/*
	Retrieve the index in Hosts array by Ip address.
	Argument:
		ipstr string
	Return value:
		-1 if note found
		index value otherwise
*/
func getHostIndexByIP( ipstr string ) int{
	for i:=0 ;i<len(Hosts); i++{
		if( ipstr == Hosts[i].IP_addr ){
			return i
		}
	}
	return -1
}

/*
	Would initialize the package-wid global variable "Hosts" 
*/
func initHostInformation(mode RunMode_t){
	var hostStr string
	switch mode {
		case mode_local:
			hostStr = "localTest.json"
		case mode_remote:
			hostStr = "servers.json"
		default:
			fmt.Println("invalid RUN MODE")
			os.Exit(1)
	} 
	UTILS_currRunMode = mode
	jData, err := ioutil.ReadFile(hostStr)
>>>>>>> d7006056134c6d97536a6a1513939b3685343f06
	exitOnErr(err, "cannot read json file:")

	hosts := []Host{}
	err = json.Unmarshal( jData, &hosts)
	exitOnErr(err, "cannot Unmarshal json file:")

	// initialize data in hosts
	for i:=0 ; i<len(hosts); i++{
		h := &(hosts[i])

		// connection
		h.Conn = nil

		// get remote ip addrs
		str,_ := net.LookupHost(h.Domain_name)
		//fmt.Println(str[0])
		h.IP_addr = string(str[0])
	}
	Hosts = hosts
}

func getLocalServers() []string{
	return []string{
		"localhost:8787",
		"localhost:9453",
	}
}

// B cast
