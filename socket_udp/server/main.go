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
	cn,err := net.ListenUDP("udp", udpAdd)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		handleMsg(cn)
	}
}

func handleMsg(cn  *net.UDPConn) {
	 var buf [512]byte 
	_,addr,err := cn.ReadFromUDP(buf[0:])
	
	if err != nil {
		fmt.Println(err)
		return
	}
	str := fmt.Sprintf("remote addr: %v  mag: %s", addr, string(buf[0:]))
	fmt.Println(str)

	daytime := time.Now().String()
	cn.WriteToUDP([]byte(daytime),addr)
}