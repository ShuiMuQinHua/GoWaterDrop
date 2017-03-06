package main

import (
	"errors"
	"fmt"
	"net"
	"net/rpc"         //实现了最基本的rpc调用
	"net/rpc/jsonrpc" //JSON-RPC协议，即实现了net/rpc包的ClientCodec接口与ServerCodec，增加了对json数据的序列化与反序列化。
	"os"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func main() {
	arith := new(Arith)
	rpc.Register(arith)

	tcpAddr, err := net.ResolveTCPAddr("tcp", ":1234")
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		//		println(conn)
		//		getByte := make([]byte, 1024)
		//		n, _ := conn.Read(getByte)
		//		println(getByte[:n])
		go jsonrpc.ServeConn(conn) //ServeConn在单个连接上执行DefaultServer
	}

}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
