package main

import (
	"fmt"
	"reflect"
)

type NotknownType struct {
	s1, s2, s3 string
}

func (n NotknownType) Stringsss() string {
	return n.s1 + "-" + n.s2 + "-" + n.s3
}

var secret interface{} = NotknownType{"Ada", "Go", "Oberon"}

func main() {
	value := reflect.ValueOf(secret)
	typ := reflect.TypeOf(secret)
	fmt.Println(typ)
	knd := value.Kind()
	fmt.Println(knd)

	for i := 0; i < value.NumField(); i++ {
		fmt.Printf("field %d:%v\n", i, value.Field(i))
	}
	
	results := value.Method(0).Call(nil)
	fmt.Println(results)
	
    for i:=0;i<typ.NumMethod() ;i++  {
        fmt.Println(typ.Method(i).Name)
	}
}
