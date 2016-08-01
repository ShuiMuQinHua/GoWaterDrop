package main

import (
	"fmt"
	"strings"
)

func main() {
	tt := "ab,cT,EC,T"
	fmt.Println(strings.ToUpper(tt))
	fmt.Println(strings.Split(tt, ","))
}
