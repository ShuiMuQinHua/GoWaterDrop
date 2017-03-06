package main

import (
	"fmt"
	"reflect"
)

type TagType struct {
	field1 bool   "an import answer"
	field2 string "the name of thing"
	field3 int    "how much there are"
}

func main() {
	tt := TagType{true, "bark obama", 1}
	for i := 0; i < 3; i++ {
		refTag(tt, i)
	}
}

func refTag(tt TagType, ix int) {
	ttType := reflect.TypeOf(tt)
	ixField := ttType.Field(ix)
	fmt.Printf("%v\n", ixField.Tag)
}
