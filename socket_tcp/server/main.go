package main

import (
	"fmt"
	"net"
)


func main(){
	add := "127.0.0.1:9090"
	ln,err := net.Listen("tcp", add)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		cn,err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleMsg(cn)
	}
}


func handleMsg(cn net.Conn) {
	remote := cn.RemoteAddr().String()
	fmt.Println(remote)
	receive := make([]byte,10)
	cn.Read(receive)
	fmt.Println("receive data:" + string(receive))
	msg := "service return"
	cn.Write([]byte(msg))
}


