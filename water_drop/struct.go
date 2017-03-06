package main

import "fmt"

type person struct {
	name string
	age  int
	Sex  bool
}

func main() {
	a := person{name: "张三", age: 12, Sex: true}
	b := person{"李四", 20, false}
	fmt.Println(a)
	a.study()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(a.name)
	fmt.Println(b.Sex)
}

func (p *person) study() {
	p.name = "王五"

}
