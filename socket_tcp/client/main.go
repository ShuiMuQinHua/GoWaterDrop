package main

import (
	"fmt"
	"net"
)

func main () {
	add := "127.0.0.1:9090"
	cn,err := net.Dial("tcp",add)
	if err != nil {
		fmt.Println(err)
	}

	msg := "client test"
	cn.Write([]byte(msg))

	recv := make([]byte,10)
	cn.Read(recv)
	fmt.Println("receive from service: " + string(recv))
}