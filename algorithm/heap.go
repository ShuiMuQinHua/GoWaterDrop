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
	/* 堆中插入数据 */
	heapInsert()

	/* 删除堆顶元素 */
	delHeapTop()

	/* 堆排序 */
	heapSort([]int{1,3,4,76,34,1,4})

	/* 用堆实现 10亿个搜索日志 如何快速获取到top10最热门搜索关键词*/

	/* 堆的应用1:优先级队列 优先级高的优先出队
		1.合并有序小文件
		2.高性能定时器
	*/
	/* 堆的应用1:求top K 
		1.静态数据取前K大
		2.动态数据取前K大
	*/
	/* 堆的应用1:求中位数 
		
	*/

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

/* 堆排序 是原地排序 整体的时间复杂度是 O(n*logn) 不是稳定的排序(因为每次都是把堆顶的元素和最后一个元素交换)
	1.建堆  时间复杂度  O(n)
	2.排序  时间复杂度  O(n*logn)   每次都把堆顶的元素和堆的最后一个元素做交换
   和快排相比的弱势：
	1.数据访问不连续。数据访问是  i 2i 2i+1 对cpu的缓存不友好
	2.同样的数据，排序过程中，交换的次数更多
*/
func heapSort(a []int){
	buildHeap(a)
	k := len(a)
	for ; k > 1; {
		a[1],a[k] = a[k],a[1]
		k--
		heapify(a,k,1)  //堆化剩下的 n-1个数据
	}
}

/* 建堆 */
func buildHeap(sli []int) {
	n := len(sli)
	for i:= n/2; i>=1; i-- {
		heapify(sli, n, i-1)
	}
}

/* 堆化 */
func heapify(sli []int, n int, i int){
	for ; ; {
		maxPos := i
		if 2*i <= n && sli[2*i]>sli[i] {
			maxPos = 2 * i
		}
		if (2*i + 1) <= n && sli[2*i + 1] > sli[maxPos] {
			maxPos = 2 * i + 1
		}
		
		if maxPos == i {
			break
		}

		sli[i], sli[maxPos] = sli[maxPos], sli[i]
		i = maxPos
	}

}


/* 获取最热门top10 */
func heapTop10(sli []int) {
	/* 
		1.数据堆化，转换成大顶堆
		2.取堆顶元素，并删除堆顶
		3.重复1 2 的操作，直到取够
	*/
	
}




func getSlice() []int {
	sli := []int{0,90,34,67,36,32,67,120,45,78,39}
	return sli
}


