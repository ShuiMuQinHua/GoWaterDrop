package main

import (
    "fmt"
)

//申明一个结构体 stockPosition
type stockPosition struct{
    ticker string
    sharePrice float32
    count float32
}

//结构体stockPosition实现了方法getValue()
func (s stockPosition) getValue() float32{
    return s.sharePrice*s.count
}

//申明一个结构体 car
type car struct{
    make string
    model string
    price float32
}

//结构体car实现了方法 getValue()
func (c car) getValue() float32{
    return c.price
}

//申明一个接口 valuable 包含getValue()方法
type valuable  interface{
    getValue() float32
}

//showValue方法 参数为一个接口valuable类型的变量
func showValue(asset valuable){
    fmt.Println(asset.getValue())
}



func main(){
    //申明一个接口类型的变量o 并给他赋值一个stockPosition结构体类型的变量
    var o valuable=stockPosition{"GOOG",577.20,4}
    //调用showValue 展示数据
    showValue(o)
    //给接口类型的变量o重新赋值，此时赋值car结构体类型的变量
    o = car{"BMW","M3",66500}

    //把接口类型的变量转换为car结构体类型 存放到v中 这是接口的类型断言
    if v, ok := o.(car); ok { // checked type assertion
        fmt.Println(v)
    }
    //varI is not of type T
    //调用showValue 展示数据
    showValue(o)
    o = stockPosition{"GOOG",577.20,4}
    //t中存取的是这个接口变量的值 如果不需要接口变量的值 只需要判断类型 就不需要赋值
    switch t:=o.(type) {
    case car:
        fmt.Println(t)
    case stockPosition: 
        fmt.Println(t) 
    case nil:
        fmt.Println(t) 
    default:
        fmt.Println(t) 
    }

    
}
