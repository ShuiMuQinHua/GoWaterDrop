package main

import "fmt"
import "os"

func main() {
	const b = 3 / 2
	const a = `0.63519283478325730257
73547328325673456348767348` //Go中多行的连接
	fmt.Println(a)
	fmt.Println(b)

	//os包获取环境变量
	var goos string = os.Getenv("GOOS")
	fmt.Printf("the system is %s\n", goos)
	path := os.Getenv("PATH")
	fmt.Printf("the path is %s\n", path)
}
