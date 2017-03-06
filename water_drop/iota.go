package main

import(
    "fmt"
)

const (
	version = "0.1.0"
	//枚举
	printReqHeader uint8 = 1 << (iota-1) //1
	printReqBody //2
	printRespHeader //4
	printRespBody //8
)

func main(){
    fmt.Println(printReqHeader)
    fmt.Println(printReqBody)
    fmt.Println(printRespHeader)
    fmt.Println(printRespBody)
}