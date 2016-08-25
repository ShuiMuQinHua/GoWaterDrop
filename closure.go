package main

import (
	"fmt"
)

//闭包,匿名函数
func main() {
	/*
	*直接调用匿名函数，func后面必须紧跟着小括号，花括号{}覆盖函数体,最后一对括号表示对该匿名函数的调用
	 */
	sum := 0
	func() {
		for i := 0; i <= 100; i++ {
			sum += i
		}
	}()
	fmt.Println(sum)

	/*
	*将匿名函数赋值给变量，然后通过变量进行调用。其实是把匿名函数的地址赋值给变量
	 */
	f()

	/*
		匿名函数和defer搭配使用
	*/
	fmt.Println(h())

	/*
		将匿名函数作为返回值
	*/
	p2 := Add2()
	fmt.Println(p2(3))
	p3 := Adder(2)
	fmt.Println(p3(3))

	var f = Adder1()
	fmt.Print(f(1), "-")
	fmt.Print(f(20), "-")
	fmt.Print(f(300))
	fmt.Println("test")

	ff := Adder2()
	for i := 0; i < 10; i++ {
		fmt.Println(ff())
	}

}

/*
将匿名函数赋值给变量，然后通过变量进行调用。其实是把匿名函数的地址赋值给变量
匿名函数可以被赋值给变量并作为值使用
*/
func f() {
	for i := 0; i < 4; i++ {
		g := func(i int) { fmt.Printf("%d", i) } //变量g其实就是匿名函数的地址
		g(i)                                     //通过变量g对匿名函数进行调用
		fmt.Printf("- g is of type %T and has value %v\n", g, g)
	}
}

/*
匿名函数和defer搭配使用.defer用于在函数返回之后，处理一些信息的。用于改变函数的返回值
函数返回1即为ret=1 又defer语句，所以在返回之后，ret++,所以最终的返回值为2
可用于在返回语句之后修改返回的error时使用。给匿名函数传参数
*/
func h() (ret int) {
	defer func() {
		ret++
	}()
	return 1
}

/*
	将匿名函数作为返回值
*/
func Add2() func(b int) int {
	return func(b int) int {
		return b + 2
	}
}

func Adder(a int) func(b int) int {
	return func(b int) int {
		return a + b
	}
}

/*
	在多次调用中，变量x的值是被保留的。闭包函数保存并积累其中的变量的值，不管外部函数退出与否，他都能够继续操作外部函数中的局部变量
	在闭包中使用到的变量可以是在闭包函数体内声明的，也可以是外部函数声明的
*/
func Adder1() func(int) int {
	var x int
	return func(delta int) int {
		x += delta
		return x
	}
}

//用闭包实现的斐波那契数列
func Adder2() func() int {
	x := 0
	y := 1
	t := 0
	i := 0
	return func() int {
		if i == 0 {
			i++
			return 0
		}
		if i == 1 {
			i++
			return 1
		}
		t = y
		y = x + y
		x = t
		i++
		return y
	}
}
