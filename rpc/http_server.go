package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"net/rpc" //rpc包实现了最基本的rpc调用
)

type Watcher int

func (w *Watcher) GetInfo(arg int, result *int) error {
	*result = 1
	return nil
}

func main() {

	http.HandleFunc("/ghj1976", Ghj1976Test)

	watcher := new(Watcher)
	rpc.Register(watcher) //注册RPC服务
	rpc.HandleHTTP()      //用于指定 RPC 的传输协议, 这里是采用 http 协议作为RPC调用的载体

	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println("监听失败，端口可能已经被占用")
	}
	fmt.Println("正在监听1234端口")
	http.Serve(l, nil)
}

func Ghj1976Test(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<html><body>ghj1976-123</body></html>")
}
