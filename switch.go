package main

import (
	"fmt"
)

func main() {
	num := 100

	switch num {
	case 98, 99:
		fmt.Println("等于98或者99")
	case 100:
		fmt.Println("等于100")
	default:
		fmt.Println("默认执行")
	}

	switch {
	case num < 100:
		fmt.Println("num比100小")
	case num > 100:
		fmt.Println("num比100大")
	case num == 100:
		fmt.Println("num等于100")
	default:
		fmt.Println("执行默认的")
	}
}
