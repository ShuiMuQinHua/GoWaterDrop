package main

import (
	"fmt"
	"os"
)

func main() {
	arg_num := len(os.Args)
	fmt.Printf("the number of input is %d\n", arg_num)

	fmt.Printf("they are: \n")
	for i := 0; i < arg_num; i++ {
		fmt.Println(os.Args[i])
	}
}
