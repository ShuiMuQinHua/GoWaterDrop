package main

import (
	"fmt"
	"os"
)

func main() {
	pid := os.Getpid()
	ppid := os.Getppid()
	//测试的方法(1)go install 会在bin目录下生成os.exe  (2)cd 到bin目录 os a b c 即可看到参数被打印出来
	args := os.Args //接收命令行参数
	fmt.Println(args)
	fmt.Println(pid)
	fmt.Println(ppid)
}
