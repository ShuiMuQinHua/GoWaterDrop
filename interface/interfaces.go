package main

import (
    "fmt"
)

type Shaper interface{
    Area() float32
    //P() float32
}

type Square struct{
    side float32
}

func (sq *Square) Area() float32{
    return sq.side*sq.side
}

func main(){
    sq1:=new(Square)
    sq1.side=5
    //多态 的 Go 版本
    var areaIntf Shaper
    //因为类型Square实现了接口Shaper 所以可以把一个Square类型的变量 赋值给Shaper接口类型的变量
    //这个时候 areaIntf接口方法指针表中的指针  就指向了Square.Area()方法
    //现在接口变量包含一个指向 Square 变量的引用，通过它可以调用 Square 上的方法 Area()
    //接口变量里包含了接收者实例的值和指向对应方法表的指针
    areaIntf = sq1  
    fmt.Println(areaIntf.Area())
}