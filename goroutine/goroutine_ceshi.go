package main

import (
	"fmt"
	"time"
)

//此方法是对管道进行读取
func c(i chan int) {
	fmt.Println(<-i)
}

func main() {
	out := make(chan int)
	//	//对管道进行发送操作
	//	out <- 2
	//	go c(out) //会照成死锁

	//在发送操作之前进行管道读取操作
	//注意的地方是：作为发送方，在准备读取之前管道是堵塞的。
	go c(out)
	//对管道进行发送操作
	out <- 2
}
