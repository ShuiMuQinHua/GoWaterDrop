package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	//循环输出一个 半金字塔
	for j := 1; j < 11; j++ {
		for k := 0; k < j; k++ {
			fmt.Print("G")
		}
		fmt.Println()
	}

	str := "GGGGGGGGGG"
	for i := 0; i < 10; i++ {
		fmt.Println(Substr(str, 0, i+1))
	}
}

//截取字符串的函数
func Substr(str string, start int, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}

	return string(rs[start:end])
}

//截取字符串的函数
func Substr1(str string, start int, end int) string {
	rs := []rune(str)
	length := len(rs)

	if start < 0 || start > length {
		panic("start is wrong")
	}

	if end < 0 || end > length {
		panic("end is wrong")
	}

	return string(rs[start:end])
}
