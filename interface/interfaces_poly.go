//接口实现的多态
package main

import (
    "fmt"
)

//申明了一个接口 Shaper 含有方法Area
type Shaper interface{
    Area() float32
}

//申明了一个结构体类型 Square
type Square struct{
    side float32
}

//结构体类型Square实现了方法Area 调用这个方法的时候 需要使用指向Square变量的指针来调用
func (sq *Square) Area() float32{
    return sq.side*sq.side
}

//申明一个结构体类型 Rectangle
type Rectangle struct{
    length,width float32
}

//结构体类型Rectangle实现了方法 Area 调用这个方法的时候 使用一个Rectangle类型的变量调用即可
func (r Rectangle) Area() float32{
    return r.length*r.width
}

func main(){
    //定义一个Rectangle类型的变量,并赋值
    r:=Rectangle{5,3}
    //&Square{5} 取Square类型的变量的地址 赋值给q
    //&符号的意思是对变量取地址
    q:=&Square{5}

    //定义一个Shaper接口类型的切片，并把上面的两个变量，作为元素赋给这个切片
    //赋值的时候  这个切片中的每一个接口元素  都有了指向各个不同类型的的结构体方法的指针
    shapes:=[]Shaper{r,q}
    //循环这个切片
    for n,_:=range shapes{
        fmt.Println(shapes[n])
        fmt.Println(shapes[n].Area())
    }
}