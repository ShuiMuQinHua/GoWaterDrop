package main

import (
	"fmt"
)

//map的key和value对调。但是要保证value值得唯一性
var (
	barval = map[string]int{"al": 34, "br": 56, "ch": 23, "de": 87}
)

func main() {
	invmap := make(map[int]string, len(barval))
	for k, v := range barval {
		invmap[v] = k
	}
	fmt.Println("inverted:")
	for k, v := range invmap {
		fmt.Printf("Key:%v,Value:%v \n", k, v)
	}
	fmt.Println()
}
