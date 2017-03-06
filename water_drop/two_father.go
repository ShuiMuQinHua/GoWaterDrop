
//多重继承
package main

import(
    "fmt"
)

type Phone struct{}

func (p *Phone) Call()string{
    return "ring ring"
}

type Caramer struct{}

func (c *Caramer) TakeAPicture()string{
    return "click"
}

type CaramerPhone struct{
    Caramer
    Phone
}

func main(){
    cp:=new(CaramerPhone)
    fmt.Println(cp.Call())
    fmt.Println(cp.TakeAPicture())
}



