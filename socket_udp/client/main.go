package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	add := "127.0.0.1:9091"
	udpAdd,err := net.ResolveUDPAddr("udp4", add)
	if err != nil {
		fmt.Println(err)
		return
	}
	conn,err := net.DialUDP("udp",nil,udpAdd)
	if err != nil {
		fmt.Println(err)
	}

	for {
		_,err = conn.Write([]byte("anything"))
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("send string anything")

		var buf [512]byte
		n,_ := conn.Read(buf[0:])
		fmt.Println(string(buf[0:n]))

		time.Sleep(3 * time.Second)
	}
	
}