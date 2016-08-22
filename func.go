package main

import "fmt"

func main() {
	a := 10
	b := 20

	fmt.Println(sum(a, b))
	fmt.Println(area(a, b))

}

func sum(a, b int) int {
	return a + b
}

func area(a, b int) int {
	return a * b
}
