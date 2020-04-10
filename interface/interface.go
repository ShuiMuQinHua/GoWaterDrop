package main

import (
	"fmt"
	"math"
)

//Interface定义了一个或一组method(s)，这些method(s)只有函数签名，没有具体的实现代码
type phic interface {
	area() float64
	perm() float64
}

type rect struct {
	hight float64
	wight float64
}

type circle struct {
	r float64
}

func (r rect) area() float64 {
	return r.hight * r.wight
}

func (r rect) perm() float64 {
	return 2 * (r.hight + r.wight)
}

func (r rect) getall() (float64, float64) {
	return r.area(), r.perm()
}

func (c circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (c circle) perm() float64 {
	return 2 * math.Pi * c.r
}

func measure(p phic) {
	fmt.Println(p)
	fmt.Println(p.area())
	fmt.Println(p.perm())
}

func main() {
	r := rect{10, 20}
	c := circle{10}
	measure(r)
	measure(c)
	fmt.Println(r.getall())
}
