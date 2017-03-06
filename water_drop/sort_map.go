package main

import (
	"fmt"
	"sort"
)

//map是无法排序的 如果想要对map排序 只能将key（或者value）拷贝到一个切片，然后对切片排序
var (
	barval = map[string]int{"al": 34, "br": 56, "ch": 23, "de": 87}
)

func main() {
	fmt.Println("unsorted:")
	for k, v := range barval {
		fmt.Printf("Key:%v,Value:%v \n", k, v)
	}
	keys := make([]string, len(barval))
	i := 0
	for k, _ := range barval {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	fmt.Println()
	fmt.Println("sorted:")
	for _, k := range keys {
		fmt.Printf("Key:%v,Value:%v\n", k, barval[k])
	}
}
