package main

import(
    "fmt"
)

//因为这个地方的ch3和ch4都只是申明了，而没有初始换 为nil的 所以case语句中 针对这两个通道的 
//发送操作和接收操作都是阻塞的
//在执行select语句的时候，系统会自上而下的判断，每一个case中的接收和发送操作是否会立即执行，
//如果可以立即执行，就执行相应的case语句，其他的case语句就会被忽略
var ch3 chan int
var ch4 chan int
var chs = []chan int{ch3,ch4}
var numbers = []int{1,2,3,4,5}

func main(){
    select{
        case getChan(0) <- getNumber(2):
            fmt.Println("1 selected")
        case getChan(1) <- getNumber(3):
            fmt.Println("2 selected")
        default:
            fmt.Println("default")

    }
}

func getNumber(i int) int {
    fmt.Printf("numbers[%d]\n",i)
    return numbers[i]
}

func getChan(i int) chan int{
    fmt.Printf("chs[%d]\n",i)
    return chs[i]
}