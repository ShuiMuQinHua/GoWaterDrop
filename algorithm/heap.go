package main

/* 堆是一种特殊的树，必须满足的条件 
	1.堆是一个完全二叉树，所以最好的存储方式是用数组来存
	2.堆中每一个节点的值都必须大于等于（或小于等于）其子树中每个节点的值
		1)每个节点的值都大于等于子树中每个节点值的堆，我们叫作"大顶堆"
		2)每个节点的值都小于等于子树中每个节点值的堆，我们叫作"小顶堆"
*/

import (
	"fmt"
)

func main(){
	fmt.Println("")
}


/* 往堆中插入数据思路  时间复杂度  O(logn)
	1.把数据插入数组的末尾
	2.就是顺着节点所在的路径，向上或者向下，对比，然后交换
	3.直到父子节点满足大小关系
*/
func heapInsert(){

}

/* 删除堆顶元素 时间复杂度  O(logn)
	1.我们把最后一个元素放到堆顶(避免空洞 一旦空洞就不满足完全二叉树的定义了)
	2.顺着路径对比，直到满足父子节点之间的大小关系
*/
func delHeapTop(){

}

/* 堆排序 */
func heapSort(a []int){
	l := len(a)
	start_node := 0 
	if l%2 == 0 {
		start_node = l/2 - 1
	} else {
		start_node := (l-1)/2 -1
	}

	for ; a[start_node] > a[start_node / 2] ; {
		
	}

	
}


func getSlice() []int {
	sli := []int{0,90,34,67,36,32,67,120,45,78,39}
	return sli
}


