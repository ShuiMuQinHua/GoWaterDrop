package main

import "fmt"

func adder() func(int) int {
    sum := 0
    return func(x int) int {
        sum += x
        return sum
    }
}

func main() {

	// pos, neg := adder(), adder()
	pos := adder()    //注意这两种方式的差异，golang的返回即复制，上面的这种赋值方式，pos和neg是不同的内存块，有完全不同的sum参数
	neg := pos
    for i := 0; i < 10; i++ {
        fmt.Println(
            pos(i),
            neg(-2*i),
        )
    }
}