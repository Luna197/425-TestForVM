package main

import (
	"fmt"
	"net"
	"os"
	"encoding/json"
	"io/ioutil"
)

type Host struct{
	Id			string	`json:"id,omitempty"`
	Domain_name	string	`json:"dname,omitempty"`
	IP_addr		string	`json:"ip,omitempty"`
	Port		string	`json:"port,omitempty"`
}

func (h Host) String() string{
	return fmt.Sprintf("<Host id:%v, dame:%v, ip:%15v, port:%v >", h.Id, h.Domain_name, h.IP_addr, h.Port)
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

func getRemoteServers() []Host{
	
	jData, err := ioutil.ReadFile("servers.json")
	exitOnErr(err, "cannot read json file:")

	hosts := []Host{}
	err = json.Unmarshal( jData, &hosts)
	exitOnErr(err, "cannot Unmarshal json file:")

	// get remote ip addrs
	for _,h:= range hosts{
		str,_ := net.LookupHost(h.Domain_name)
		h.IP_addr = string(str[0])
	}
	return hosts
}

func getLocalServers() []string{
	return []string{
		"localhost:8787",
		"localhost:9453",
	}
}

// B cast
