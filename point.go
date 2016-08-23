package main

import (
	"fmt"
)

func main() {
	a := 5          //注意  你不能得到一个文字或者常量的地址
	fmt.Println(&a) //返回的变量a的内存地址  实际输出的地址 根据机器的不同而不同，同样的代码两次执行的输出结果不同，因为机器之后的内存布局不同，另外在程序执行时，每次都是动态的分配地址空间

	//	var intP *int
	//	intP = &a //此时intP指向a

	//这个表达式是恒为真的
	if a == *(&a) {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}

}
