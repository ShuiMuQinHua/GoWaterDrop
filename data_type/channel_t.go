/*
1. 通道的方向
2. 与select结合
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	/* 通道方向 
		ch1 := make(chan int)
		go pump(ch1)
		go suck(ch1)
		time.Sleep(1e9)
		var c = make(chan int) // 通道在创建的时候都是双向的
		go source(c)
		go sink(c)
	*/

	/* 结合select使用 
		f := make(chan int, 1)
		timeout := time.After(time.Second * 2) //
		t1 := time.NewTimer(time.Second * 3)   // 效果相同 只执行一次
		go  pushInt(f)
		
		for {
			select {
			case cd := <-f :
				fmt.Println("get data from channel : %v", cd)
			case <-time.After(3*time.Second):
				fmt.Println("timeOut")
			case <-t1.C:      // 代码段3
				fmt.Println("3s定时任务")
			case <-timeout:   // 代码段4
				fmt.Println("2s定时输出")
			default:          // 代码段5
				fmt.Println("default")
				time.Sleep(time.Second * 1)
			}
		}
	*/

	/* 利用通道，计算素数 ???
		ch := make(chan int) // Create a new channel.
		go generate(ch) // Start generate() as a goroutine.
		for {
			prime,open := <-ch
			if !open{
				break
			}
			fmt.Print(prime, " ")
			ch1 := make(chan int)
			go filter(ch, ch1, prime)
			ch = ch1
		}
	*/
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


func pushInt(ch chan int) {
	time.Sleep(4*time.Second)
	ch <- 1
	i := 0
	for {
		if i == 10 {
			close(ch)  //读取关闭后的无缓存通道，不管通道中是否有数据，返回值都为0和false
			break
		}
		ch <- 1
		i++
	}
}

//素数:除了一和它自身以外，没有其它的因数
func generate(ch chan int) {
    for i := 2;i<=1000 ; i++ {
        ch <- i // Send 'i' to channel 'ch'.
    }
    defer close(ch)
}

func filter(in, out chan int, prime int) {
    for {
        i := <-in // Receive value of new variable 'i' from 'in'.
        if i%prime != 0 {
            out <- i // Send 'i' to channel 'out'.
        }
    }
}
