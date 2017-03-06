//使用聚合的方法给 在类型中嵌入功能

package main

import(
    "fmt"
)

type Log struct{
    msg string
}

type Customer struct{
    Name string
    log *Log
}

func (l *Log) Add(s string){
    l.msg+="\n"+s
}

func (l *Log) String() string{
    return l.msg
}

func (c *Customer) Log() *Log{
    return c.log
}

func main(){
    c:=new(Customer)
    c.Name = "bbbb"
    c.log=new(Log)

    c.log.msg="11111"
    c.Log().Add("22222")

    fmt.Println(c.Log())
}

