package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:1234") //客户端先用rpc.DialHTTP和RPC服务器进行一个链接(协议必须匹配).
	if err != nil {
		fmt.Println("链接rpc服务器失败:", err)
	}
	var reply int
	err = client.Call("Watcher.GetInfo", 1, &reply) //然后通过返回的client对象进行远程函数调用. 函数的名字是由client.Call 第一个参数指定(是一个字符串).
	if err != nil {
		fmt.Println("调用远程服务失败", err)
	}
	fmt.Println("远程服务返回结果：", reply)
}
