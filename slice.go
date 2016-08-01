package main

import "fmt"

type phone struct {
	price  int
	color  string
	number int
}

func main() {
	a := 2
	b := "wjdiwj232"
	c := [3]string{"qweq", "wqrw", "dvnsdj"}
	d := []int{2, 4, 7, 9, 3}
	e := phone{price: 123, color: "red", number: 123456}
	f := map[int]int{1: 2, 2: 5, 3: 7, 9: 8, 100: 68}

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}
