package main

import (
    "google.golang.org/grpc"

    "fmt"
    "context"
    "time"
    
    
)

import (
    "myprotos/customer"
)

const (
    ADDRESS = "localhost:50051"
)


func main(){
    //通过grpc 库 建立一个连接
    conn ,err := grpc.Dial(ADDRESS,grpc.WithInsecure())
    if err != nil{
        return
    }
    defer conn.Close()
    //通过刚刚的连接 生成一个client对象。
    c := customer.NewCustomerClient(conn)
    
    reqstreamData := &customer.CustomerRequest{}
    res,_ := c.CreateCustomer(context.Background())
    i:=0
    for {
        err := res.Send(reqstreamData)
        if err != nil {
            fmt.Println(err.Error())
        }
        fmt.Println(i)
        i++
        if i > 10 {
            break
        }
    }

    // _,err = res.CloseAndRecv()
    // if err != nil {
    //     fmt.Println(err.Error())
    // }
    
    // fmt.Println("closed")

    time.Sleep(3 * time.Second)

    for {
        err := res.Send(reqstreamData)
        if err != nil {
            fmt.Println(err.Error())
        }
        fmt.Println(i)
        i++
        if i > 20 {
            break
        }
    }

    // _,err = res.CloseAndRecv()
    // if err != nil {
    //     fmt.Println("err: " + err.Error())
    // }
    
    // fmt.Println("closed")
    time.Sleep(3 * time.Second)

    for {
        err := res.Send(reqstreamData)
        if err != nil {
            fmt.Println(err.Error())
        }
        fmt.Println(i)
        i++
        if i > 30 {
            break
        }
    }

    res.CloseAndRecv()



    for {
        err := res.Send(reqstreamData)
        if err != nil {
            fmt.Println(err.Error())
        }
        fmt.Println(i)
        i++
        if i > 40 {
            break
        }
    }
    res.CloseAndRecv()

    select {
    }

}