package main

import (
	//	"fmt"
	"time"
)

func main() {
	//	ticker := time.NewTimer(time.Second * 5)
	//	go func() {
	//		for _ = range ticker.C { //
	//			fmt.Println("aa")
	//		}
	//	}()
	msg1 := make(chan interface{})
	demo(msg1)
}

//实现了定时任务
func demo(input chan interface{}) {
	t1 := time.NewTimer(time.Second * 5)
	t2 := time.NewTimer(time.Second * 10)
	msg := make(chan interface{})

	for {
		select {
		case msg <- input:
			println(msg)

		case <-t1.C:
			println("5s timer")
			t1.Reset(time.Second * 5)

		case <-t2.C:
			println("10s timer")
			t2.Reset(time.Second * 10) //重置一个10s的定时器
		}
	}
}
