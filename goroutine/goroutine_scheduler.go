package main

import(
	"fmt"
	"time"
)

func main(){
	go func(){
		for{

		}
	}()

	go func(){
		for{
			time.Sleep(1 * time.Second)
			fmt.Println(1)
		}
	}()

	select {

	}
}