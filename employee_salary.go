package main

import "fmt"

type Employee struct {
	salary int
}

func main() {
	e := new(Employee)
	e.salary = 2000

	fmt.Println(e.giveRaize())
}

func (e *Employee) giveRaize() int {
	return e.salary * 2
}
