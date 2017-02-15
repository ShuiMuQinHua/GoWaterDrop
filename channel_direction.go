package main

//通道的方向
import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	go pump(ch1)
	go suck(ch1)
	time.Sleep(1e9)
	var c = make(chan int) // 通道在创建的时候都是双向的
	go source(c)
	go sink(c)
}

func pump(ch chan int) {
	for i := 0; ; i++ {
		ch <- i
	}
}

func suck(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}

//通道在创建的时候，都是双向的，但是我们可以分配有方向的通道变量
func source(ch chan<- int) {
	//只能向通道中写入数据
	for {
		ch <- 1
	}
}
func sink(ch <-chan int) {
	//只能从通道中读取数据
	for {
		<-ch
	}
}
