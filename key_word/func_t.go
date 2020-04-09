package main

import (
	"fmt"
)

func main(){
	/* 传递变长参数
		min := Min(1,3,100,0,300)
		fmt.Println(min)
		s := []int{3,4,1,8}
		m := Min(s...)
		fmt.Println(m)
		return
	*/

	/* 函数作为参数传递 */
	callback(1, Add)
}


/*
	传递变长参数
*/
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

func callback(y int, f func(int, int)) {
	f(y, 2)
}

func Add(a, b int) {
	fmt.Printf("the sum of %d and %d is: %d\n", a, b, a+b)
}
