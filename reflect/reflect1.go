// package main

// import (
// 	"fmt"
// 	"reflect"
// )

// func main() {
// 	var x float64 = 3.4
// 	fmt.Println("type:", reflect.TypeOf(x))
// 	v := reflect.ValueOf(x)
// 	fmt.Println("value:", v)
// 	fmt.Println("type", v.Type())
// 	fmt.Println("kind:", v.Kind())
// 	fmt.Printf("value is %5.2e\n", v.Interface())
// 	y := v.Interface().(float64)
// 	fmt.Println(y)
// }

package main

import(
	"fmt"
)

type Persion struct{
	Name string
}

func main () {
	fmt.Println("hello https://tool.lu/")
	p := Persion{
		Name: "cyy",
	}
	checkReflect(p)
}

 
func (p *Persion) check(){
   fmt.Println("1234")
}

func checkReflect(test interface{}) {
	if _, ok := test.(Persion); ok {
		realInstance := test.(Persion)
		fmt.Println("12321")
		
		realInstance.check()
		
	}
	
}

