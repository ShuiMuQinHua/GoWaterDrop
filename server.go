package main

import(
    "fmt"
    "net"
)

func main(){
    fmt.Println("staring the server...")
    //创建listener
    listener,err := net.Listen("tcp","localhost:5000")
    if err!=nil{
        fmt.Println("error listenning",err.Error())
        return
    }
    //死循环等待客户端连接
    for{
        conn,err := listener.Accept()
        if err!=nil{
            fmt.Println("connect error",err.Error())
            return
        }
        //连接成功 对连接开启go协程
        go doServerStuff(conn)
    }
}

func doServerStuff(conn net.Conn){
    for {
        //创建一个byte(字节)类型的切片 初始化有512个空间
        buf := make([]byte, 512)
        //从连接中获取数据，存入buf中
        _, err := conn.Read(buf)
        if err != nil {
            fmt.Println("Error reading", err.Error())
            return //终止程序
        }
    //??? 切片类型转换为字符串???
    fmt.Printf("Received data: %v", string(buf))
    }
}