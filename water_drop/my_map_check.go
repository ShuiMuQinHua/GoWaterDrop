package main

import "fmt"
import "encoding/json"


type MyTest struct {
 Name string
 Age map[int]*ToTest
}

type ToTest struct {
 Age int
}
func main () {
	toTest := ToTest{Age:1}
	age := map[int]*ToTest{1:&toTest}
	mytest := MyTest{Name:"cyy",Age:age}
	res, _ := json.Marshal(mytest)
	
	fmt.Println(string(res))
}