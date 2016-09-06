package main

import(
    "fmt"
)

type Base struct{}

func (Base) Magic(){
    fmt.Println("base magic")
}

func (self Base) MoreMagic(){
    self.Magic()
    self.Magic()
}

type Voodoo struct{
    Base
}

//结果会输出3个"base magic"
func main(){
    v:=new(Voodoo)
    v.Magic()
    v.MoreMagic()
}