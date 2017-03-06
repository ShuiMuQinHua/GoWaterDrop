package main

import (
	"fmt"
)

func main() {
	var maplit map[string]int //声明了一个map类型的变量，名字是maplit,它的key类型为string，value的类型为int
	var mapassigned map[string]int
	maplit = map[string]int{"one": 1, "two": 2}

	mapassigned = maplit
	mapassigned["two"] = 3 //注意map是引用，所以现在修改了mapassigned["two"],那么maplit对应的也会改变

	mapcreate := make(map[string]float32) //使用make函数声明并赋值一个map类型的变量。map是引用类型，用make来分配变量
	mapcreate["key1"] = 4.5
	mapcreate["key2"] = 3.1315

	val1, ok := mapcreate["key1"] //key1存在，且val不为空  所以 val1=4.5 ok=true
	val3, ok := mapcreate["key3"] //key3不存在   所以val3为空且      ok=false

	if _, flag := mapcreate["key1"]; flag {
		delete(mapcreate, "key1")
	}
	fmt.Println(maplit["one"])
	fmt.Println(mapcreate["key1"])
	fmt.Println(mapcreate["key2"])
	fmt.Println(maplit["two"])
	fmt.Println(maplit["ten"])

}
