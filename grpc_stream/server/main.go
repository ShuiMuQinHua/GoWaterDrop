package main

import (
    "fmt"
    "google.golang.org/grpc"
    "net"
    "io"
)


import (
    "myprotos/customer"
)
    


const PORT  = ":50051"

type server struct {
}

//服务端 单向流
// func (s *server)GetStream(req *pro.StreamReqData, res pro.Greeter_GetStreamServer) error{
//     i:= 0
//     for{
//         i++
//         res.Send(&pro.StreamResData{Data:fmt.Sprintf("%v",time.Now().Unix())})
//         time.Sleep(1*time.Second)
//         if i >10 {
//             break
//         }
//     }
//     return nil
// }


func (this *server) CreateCustomer(stream customer.Customer_CreateCustomerServer) error {
   i := 0
    for {
		_, err := stream.Recv()
		if err == io.EOF {
			stream.SendAndClose(&customer.CustomerResponse{})
			break
		}

		if err != nil {
			//err log
			break
		}
        i++ 
    }

    fmt.Println(i)
    return nil
}

func main(){
    //监听端口
    lis,err := net.Listen("tcp",PORT)
    if err != nil{
        return
    }
    //创建一个grpc 服务器
    s := grpc.NewServer()
    //注册事件
    customer.RegisterCustomerServer(s,&server{})
    //处理链接
    s.Serve(lis)
}