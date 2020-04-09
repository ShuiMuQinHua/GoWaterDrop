package main

import (
	"fmt"
)

func main(){
	/* defer调用顺序
	   	in func1 at the top
		in func1 at the bottom
		func2:defered until the calling function end 

		fmt.Println("in func1 at the top")
		defer func2() //defer用于在被调用的函数执行完成之后，再执行
		fmt.Println("in func1 at the bottom")  
	*/

	/*defer 的执行顺序是先进后出 
		打印后
		打印中
		打印前
		panic: 触发异常
	
		defer func() {fmt.Println("打印前")}()
		defer func() {fmt.Println("打印中")}()
		defer func() {fmt.Println("打印后")}()
		panic("触发异常")
	*/

	func3()
}

func func2() {
	fmt.Println("func2:defered until the calling function end")
}


/*
	defer先进后出 先打印1 再打印0
*/
func func3(){
	i := 0
	defer fmt.Println(i)
	i++
	defer fmt.Println(i)
	return
}

