package main

/* 
	hash table hash函数的基本设计要求
	1.散列函数计算得到的散列值是一个非负整数
	2.如果 key1 = key2，那 hash(key1) == hash(key2)
	3.如果 key1 ≠ key2，那 hash(key1) ≠ hash(key2

*/

import (
	"fmt"
)

func main(){
	fmt.Println("")

	/* 散列冲突  解决办法
		1.开放寻址法   (线性探测/二次探测/双重散列)
		2.链表法(更常用) 在散列表中，每个“桶（bucket）”或者“槽（slot）”会对应一条链表，所有散列值相同的元素我们都放到相同槽位对应的链表中
	*/

	/* 用散列表 实现Word文档中单词拼写检查的功能 
		用散列表存所有的词库，然后当用户输入的时候做检测
	*/
}

