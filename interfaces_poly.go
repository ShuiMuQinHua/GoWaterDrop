//接口实现的多态
package main

import (
    "fmt"
)

type Shaper interface{
    Area() float32
}

type Square struct{
    side float32
}

func (sq *Square) Area() float32{
    return sq.side*sq.side
}

type Rectangle struct{
    length,width float32
}

func (r Rectangle) Area() float32{
    return r.length*r.width
}

func main(){
    r:=Rectangle{5,3}
    q:=&Square{5}

    shapes:=[]Shaper{r,q}
    for n,_:=range shapes{
        fmt.Println(shapes[n])
        fmt.Println(shapes[n].Area())
    }
}