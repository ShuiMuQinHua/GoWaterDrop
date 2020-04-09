package main

import (
	"fmt"
	"time"
)

type myTest struct {
	sli []string
	age int
}

var mapChan = make(chan *myTest, 10)

func main() {
	
	go func(){
		for{
			bs := <-mapChan
			bs.age  =30
			fmt.Println(bs)
		}
	}()


	go func(){
		for{
			bs := <-mapChan
			bs.age  =40
			fmt.Println(bs)
		}
	}()
	
	aa := &myTest{
		sli:[]string{},
		age:2,
	}

	for i:=0;i <100;i++ {
		mapChan <- aa
	}

	aa.age = 30

	time.Sleep(10 * time.Second)
	
}
