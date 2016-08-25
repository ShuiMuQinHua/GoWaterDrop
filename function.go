package main

import (
	"fmt"
)

func main() {
	//	func1()
	//	a()
	//f()
	//	result := 0
	//	for i := 0; i <= 10; i++ {
	//		result = fibonacci(i)
	//		fmt.Println(result)
	//	}
	callback(1, Add)
}

func func1() {
	fmt.Println("in func1 at the top")
	defer func2() //defer用于在被调用的函数执行完成之后，再执行
	fmt.Println("in func1 at the bottom")
}

func func2() {
	fmt.Println("func2:defered until the calling function end")
}

func a() {
	i := 0
	defer fmt.Println(i)
	i++
	defer fmt.Println(i)
	return
}

func f() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

//递归函数,计算斐波那契数列(递归调用可能会出现栈溢出 在go中  我们使用“懒惰求值”法解决这个问题，用channer 和goroutine来解决)
func fibonacci(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return
}

//将函数作为参数
func callback(y int, f func(int, int)) {
	f(y, 2)
}
func Add(a, b int) {
	fmt.Printf("the sum of %d and %d is: %d\n", a, b, a+b)
}
