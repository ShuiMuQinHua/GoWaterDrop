package main

import (
	"fmt"
	"os"
)

func main() {
	num1 := os.Getpid()
	num2 := os.Getppid()
	fmt.Println(num1)
	fmt.Println(num2)
}
