package main

import (
	"fmt"
)

type struct1 struct {
	i1  int
	f1  float32
	str string
}

func main() {
	mt := new(struct1)
	mt.i1 = 12
	mt.f1 = 23.23
	mt.str = "oeowoere23502"

	fmt.Println(mt.i1)
	fmt.Println(mt.f1)
	fmt.Println(mt.str)
}
