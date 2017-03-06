package main

import (
	"fmt"
)

type Integer int

type LessAdder interface {
	Less(b Integer) bool
	Add(b Integer)
}

/*
* 在GO语言中可以根据 func (a Integer) Less(b Integer) bool 自动生成
* func (a *Integer) Less(b Integer) bool {
    return (*a).Less(b)
} 这个方法  因为在得到的函数中 *a是最开始的对象 也是a本身的值
*/
func (a Integer) Less(b Integer) bool {
	return a < b
}

/*
* 在GO中无法根据 func (a *Integer) Add(b Integer) 自动生成
* func (a Integer) Add(b Integer){
    (&a).Add(b)
} 方法 在后面这个方法中 &a 是a这个变量的地址 也就是 调用的参数变了  对对象没有影响
* 所以这个是无法转换的
*/
func (a *Integer) Add(b Integer) {
	*a += b
}

func main() {
	//申明了一个Integer类型的变量 并赋值
	var a Integer = 1
	//申明一个接口类型的LessAdder类型的变量 并尝试赋值 &a 对变量取地址
	var b LessAdder = &a
	//申明一个接口类型的LessAdder类型的变量 并尝试赋值  这个不奏效
	var b LessAdder = a
}
