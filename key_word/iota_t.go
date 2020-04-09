package main

import(
    "fmt"
)

const (
	version = "0.1.0"
	//枚举
	printReqHeader uint8 = 1 << (iota-1) //1    // iota只能在const表达式中使用，且初始化为0,是一个常量计数器，且下面每增加一行  值+1
	printReqBody //2    // 1 << (1-1)
	printRespHeader //4 // 1 << (2-1)
	_    //跳过一个值
	printRespBody //16  // 1 << (4-1)
)


const (
    IgEggs int = 1 << iota // 1 << 0 which is 00000001
    IgChocolate                             // 1 << 1 which is 00000010
    IgNuts                                       // 1 << 2 which is 00000100
    IgStrawberries                        // 1 << 3 which is 00001000
    IgShellfish                                // 1 << 4 which is 00010000
)

const (
    _           = iota                   // ignore first value by assigning to blank identifier
    KB float64 = 1 << (10 * iota)       // 1 << (10*1)
    MB                                   // 1 << (10*2)
    GB                                   // 1 << (10*3)
    TB                                   // 1 << (10*4)
    PB                                   // 1 << (10*5)
    EB                                   // 1 << (10*6)
    ZB                                   // 1 << (10*7)
    YB                                   // 1 << (10*8)
)

// n << m   结果为  n * 2的m次方  注意：把 n 左移动 m 位

func main(){
    fmt.Println(printReqHeader)
    fmt.Println(printReqBody)
    fmt.Println(printRespHeader)
	fmt.Println(printRespBody)
	


	fmt.Println(IgEggs)
    fmt.Println(IgChocolate)
    fmt.Println(IgNuts)
	fmt.Println(IgStrawberries)
	fmt.Println(IgShellfish)


	fmt.Println(KB)
    fmt.Println(MB)
    fmt.Println(GB)
	fmt.Println(TB)
	fmt.Println(PB)
}