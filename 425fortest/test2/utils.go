/*
	helper functions to called in the package

	Notice:
		Must initialize before use, must follow the order is important
		1. call initHostInformation(mode RunMode_t)
			would load informations from file: `servers.json`
			Side Effect:
				All of the information would be stored in the `Hosts`.

		2. call setMyHostInformation(port string, name string)
			use this function once you know the local server's configuration
			Side Effect:
				initialize variable: `MyHostIndex`

*/

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

type RunMode_t int

const (
	mode_local  RunMode_t = 0
	mode_remote RunMode_t = 1
)

type Host struct {
	Id          string `json:"id,omitempty"`
	Domain_name string `json:"dname,omitempty"`
	IP_addr     string `json:"ip,omitempty"`
	Port        string `json:"port,omitempty"`
	UserName    string
	Conn        net.Conn
}

// use `initHostInformation()` to initialize this function
var Hosts []Host
var MyHostIndex int = -1

// indicate the mode that current running
// 	do not used in other files
var UTILS_currRunMode RunMode_t

func (h Host) String() string {
	return fmt.Sprintf("<Host id:%v, dame:%v, ip:%v, port:%v, conn:%v>", h.Id, h.Domain_name, h.IP_addr, h.Port, h.Conn)
}

func exitOnErr(err error, str string) {
	if err != nil {
		fmt.Printf("%v%v", str, err.Error())
		os.Exit(1)
	}
}

func getLocalIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	exitOnErr(err, "cannot connect to 8.8.8.8:80:")

	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	host, _, err := net.SplitHostPort(localAddr.String())
	exitOnErr(err, "Cannot split local IP:")
	return host
}

func setMyHostInformation(port string, name string) {
	if UTILS_currRunMode == mode_local {
		for i := len(Hosts) - 1; i >= 0; i-- {
			h := &(Hosts[i])
			if h.Port == port {
				MyHostIndex = i
				h.UserName = name
				return
			}
		}
	} else {
		myIP := getLocalIP()
		for i := len(Hosts) - 1; i >= 0; i-- {
			h := &(Hosts[i])
			if h.IP_addr == myIP && h.Port == port {
				MyHostIndex = i
				h.UserName = name
				return
			}
		}
	}
	var load_file string
	switch UTILS_currRunMode {
	case mode_local:
		load_file = "localTest.json"
	case mode_remote:
		load_file = "servers.json"
	default:
		fmt.Println("Please call initHostInformation() first")
		os.Exit(1)
	}
	fmt.Printf("Local Server is not listed in the `%v`", load_file)
	os.Exit(1)
}

/*
	Search the Host array and return the index of matching host
	Used in accepting a new connection
*/
func findHostIndexByConn(conn net.Conn) int {
	if conn == nil {
		return -1
	}
	remoteAddr := conn.RemoteAddr()
	rhost, rport, _ := net.SplitHostPort(remoteAddr.String())
	if UTILS_currRunMode == mode_local {
		for i := 0; i < len(Hosts); i++ {
			fmt.Printf("check port:%v, cur port: %v\n", Hosts[i].Port, rport)
			if rport == Hosts[i].Port {
				return i
			}
		}
		return -1
	} else {
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
func getHostIndexByIP(ipstr string) int {
	for i := 0; i < len(Hosts); i++ {
		if ipstr == Hosts[i].IP_addr {
			return i
		}
	}
	return -1
}

/*
	Would initialize the package-wid global variable "Hosts"
*/
func initHostInformation(mode RunMode_t) {
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
	exitOnErr(err, "cannot read json file:")

	hosts := []Host{}
	err = json.Unmarshal(jData, &hosts)
	exitOnErr(err, "cannot Unmarshal json file:")

	// initialize data in hosts
	for i := 0; i < len(hosts); i++ {
		h := &(hosts[i])

		// connection
		h.Conn = nil

		// get remote ip addrs
		str, _ := net.LookupHost(h.Domain_name)
		//fmt.Println(str[0])
		h.IP_addr = string(str[0])
	}
	Hosts = hosts
}

func getLocalServers() []string {
	return []string{
		"localhost:8787",
		"localhost:9453",
	}
}

// B cast
