package main

import (
	"fmt"
	"strconv"
	//	"strings"
)

func main() {
	//不要在多个分支里同时都使用return返回数据
	if 1 == 3 {
		fmt.Println("if")
	} else if 2 == 2 && 3 == 3 {
		fmt.Println("else if")
	} else {
		fmt.Println("else")
	}

	if a := 5; a > 4 { //这样的分支语句  多了一个初始化  前面是初始化，后面是条件
		fmt.Println("a大于4")
	}

	str := "weqw24234"
	in1, err := strconv.Atoi(str)
	if err != nil {
		fmt.Print("error")
		return
	}
	fmt.Print(in1)
}
