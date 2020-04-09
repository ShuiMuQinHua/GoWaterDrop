package main

import (
	"fmt"
	"time"
)


func main(){
	c := make(chan int, 1)
	go  pushInt(c)

	select {
	case <-c :
		fmt.Println("get data from channel")
	case <-time.After(3*time.Second):
		fmt.Println("timeOut")
	}
	
	

	// consume(c)
}


func pushInt(ch chan int) {
	time.Sleep(4*time.Second)
	ch <- 1
	// i := 0
	// for {
	// 	if i == 10 {
	// 		close(ch)
	// 		break
	// 	}
	// 	ch <- 1
	// 	i++
	// }
}


func consume(inCh <-chan int) {
    i := 0
    for {
		fmt.Printf("for: %d\n", i)
		
		// var x int
		// var open bool

        select {
        case x, open := <-inCh:
            if !open {
				goto HEAR
                // break  //break 并不能跳出 for-select结构
            }
            fmt.Printf("read: %d\n", x)
        }
		i++
		
		// if !open {
		// 	break
		// }
	}
	
	HEAR:
		fmt.Println("cycle is end")

    fmt.Println("combine-routine exit")
}