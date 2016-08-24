package main

import "fmt"

func main() {
	a := 10
	b := 20
	c := 100
	x := Min(1, 3, 2, 4)
	arr := []int{7, 9, 5, 6}
	y := Min(arr...)

	fmt.Println(sum(a, b))
	fmt.Println(area(a, b))
	fmt.Println(sum2(a, b))
	point(&c)
	fmt.Println(c)
	fmt.Println(x)
	fmt.Println(y)

}

func sum(a, b int) int {
	return a + b
}

func area(a, b int) int {
	return a * b
}

func sum2(a, b int) (c int) {
	c = a + b
	return
}

func point(d *int) {
	*d = 400
}

//传递了变长参数
func Min(a ...int) int {
	if len(a) == 0 {
		return 0
	}
	min := a[0]
	for _, v := range a {
		if v < min {
			min = v
		}
	}
	return min
}
