package main

import (
	"os"
	"fmt"
	"net"
)


func main() {
	if len(os.Args) == 3 {
		// first is ip
		strIP := os.Args[1]
		// second is port
		strPort := os.Args[2]
		// check parameter
		server := strIP + ":" + strPort
		hawkServer,err := net.ResolveTCPAddr("tcp", server)
		if err != nil {
			fmt.Printf("resolve address [%s] failed error:[%s]", server, err.Error())
			os.Exit(-1)
		}
		_,connErr := net.DialTCP("tcp",nil,hawkServer)
		if connErr != nil {
			fmt.Printf("connect to server [%s] error: [%s]", server, connErr.Error())
        	os.Exit(-1)
		}
		fmt.Printf("connect to server [%s] success", server)
	} else {
		fmt.Println("we need ip address and port!")
	}
}