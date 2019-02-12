// THis file is for testing
package main

import (
	"fmt"
)

func main(){
	
	// test local IP
	printTestString("test local ip",0)
	fmt.Printf("local IP: %v \n", getLocalIP())
	
	// test json file load successfully
	printTestString("test json file",0)
	hs := getRemoteServers()
	fmt.Printf("Total number of servers: %2v (should be 9 on the server)\n", len(hs))
	


	fmt.Println("Test complete")
}

func printTestString(str string, indent int){
	for i:=0 ; i<indent; {
		fmt.Printf("\t")
	}
	fmt.Printf("===== %v =====\n",str)
}