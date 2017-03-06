package main

import "fmt"

//var a = "G"
//func main() {
//	n()
//	m()
//	n()
//}

//func n() {
//	fmt.Println(a)
//}
//func m() {
//	a := "O" //如果是:=  则结果输出 G O G   如果是 = 则输出G O O
//	fmt.Println(a)
//}

var a string

func main() {
	a = "G"
	fmt.Println(a)
	f1()
}

func f1() {
	a := "O"
	fmt.Println(a)
	f2()
}

func f2() {
	fmt.Println(a) //输出G O G
}
