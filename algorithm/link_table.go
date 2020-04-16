package main

import (
	"fmt"
)

type Elem int

type LinkNode struct {
    Data Elem
    Next *LinkNode
}

//生成头节点
func New() *LinkNode {
    //下面的data可以用来表示链表的长度
    return &LinkNode{0, nil}
}

func main(){
	/* 1. LRU 缓存淘汰算法 最近最少使用 
		我们需要维护一个按照访问时间从大到小有序排列的链表结构。
		因为缓存大小有限，当缓存空间不够，
		需要淘汰一个数据的时候，我们就直接将链表头部的结点删除	
		当要缓存某个数据的时候，先在链表中查找这个数据。
		如果没有找到，则直接将数据放到链表的尾部；如果找到了，
		我们就把它移动到链表的尾部。因为查找数据需要遍历链表，
		所以单纯用链表实现的 LRU 缓存淘汰算法的时间复杂很高，是 O(n)
	*/

	/* 2. 单链表反转 
	nodes := getNodeList()
	printNode(nodes)
	nodes = reverseNode(nodes)
	printNode(nodes)
	*/

	/* 3. 链表两两反转 
	nodes1 := getNodeList()
	printNode(nodes1)
	nodes1 = adjacentChange(nodes1)
	printNode(nodes1)
	*/

	/* 4. 检查链表是否有环 
	nodes := getNodeList()
	check := HasCycle(nodes)
	fmt.Println(check)
	*/

	/* 5. 两个有序单链表合并 
	nodes1 := getNodeList()
	nodes2 := getNodeList()
	nodes := MergeSortedList(nodes1, nodes2)
	printNode(nodes)
	*/

	/* 6.删除链表倒数第n个元素 
	nodes := getNodeList()
	nodes = DeleteBottomN(nodes, 2)
	printNode(nodes)
	*/

	/* 7.求链表的中间结点 
	nodes := getNodeList()
	nodes = FindMiddleNode(nodes)
	printNode(nodes)
	*/

}

/* 
我的思路是这样的：我们维护一个有序单链表，越靠近链表尾部的结点是越早之前访问的。当有一个新的数据被访问时，我们从链表头开始顺序遍历链表。
 1. 如果此数据之前已经被缓存在链表中了，我们遍历得到这个数据对应的结点，并将其从原来的位置删除，然后再插入到链表的头部。
 2. 如果此数据没有在缓存链表中，又可以分为两种情况：如果此时缓存未满，则将此结点直接插入到链表的头部；如果此时缓存已满，则链表尾结点删除，将新的数据结点插入链表的头部
*/
 func LRU(){

}

/* 
	单链表反转
	链表中环的检测
	两个有序的链表合并
	删除链表倒数第 n 个结点
	求链表的中间结点
*/

/* 单链表反装 */
func reverseNode(head *LinkNode) *LinkNode {
	//  先声明两个变量
	//  前一个节点
		var preNode *LinkNode
		preNode = nil
	//  后一个节点
		nextNode := new(LinkNode)
		nextNode = nil
		for head != nil {
			//  保存头节点的下一个节点
			nextNode = head.Next
			//  将头节点指向前一个节点
			head.Next = preNode
			//  更新前一个节点
			preNode = head
			//  更新头节点
			head = nextNode
		}
		return preNode
}

/* 链表两两交换 1->2->3->4->5->6 变成 2->1->4->3->6->5  */
func adjacentChange(head *LinkNode) *LinkNode {
	i := 0

	retNode := new(LinkNode)
	retNode = nil

	pre := new(LinkNode)
	pre = head   //始终指向两两交换元素的前驱

	for head != nil && head.Next != nil {
		i++
		fNode := head    //两两交换的第一个元素
		sNode := head.Next  //两两交换的第二个元素

		pre.Next = sNode
		fNode.Next = sNode.Next
		sNode.Next = fNode

		pre = fNode
		
		head = fNode.Next

		if i == 1 {
			retNode = sNode
		}
	}
	return  retNode
}
	
/* 判断单链表是否有环 用的是快慢指针 如果有环，快指针一定可以追上慢指针，也就是  slow==fast (跑道  快的追上慢的) */
func  HasCycle(head *LinkNode) bool {
	if nil != head {
		slow := head
		fast := head
		for nil != fast && nil != fast.Next {
			slow = slow.Next
			fast = fast.Next.Next
			if slow == fast {
				return true
			}
		}
	}
	return false
}

/* 两个有序单链表合并 */
func MergeSortedList(head1, head2 *LinkNode) *LinkNode {
	if head1 == nil  {
		return head2
	}
	if head2 == nil  {
		return head1
	}
 
	l := &LinkNode{Next: &LinkNode{}}
	cur := l.Next
	l1 := head1
	l2 := head2
	for l1 != nil && l2 != nil {
		if l1.Data > l2.Data {
			cur.Next = l2
			l2 = l2.Next
		} else {
			cur.Next = l1
			l1 = l1.Next
		}
		cur = cur.Next
	}
 
	if l1 != nil {
		cur.Next = l1
	} else if l2 != nil {
		cur.Next = l2
	}
 
	return l.Next.Next
}


/*
	删除倒数第N个节点
*/
func  DeleteBottomN(head *LinkNode, n int) *LinkNode {
	if n <= 0 || head == nil  || head.Next == nil {
		return head
	}
 
	fast := head
	//循环完之后，fast指向单链表的正数第n个元素
	for i := 1; i <= n && fast != nil; i++ {
		fast = fast.Next   
	}
 
	if nil == fast {
		return head
	}
	
	/*
	slow和head一起往后走，直到链表结束。slow和fast共同走s个元素end。
	那么slow剩下的还需要走的距离，就是之前fast已经走的n，
	即当前slow指向的就是倒数第n个元素的前驱 
	*/
	slow := head
	for nil != fast.Next {
		slow = slow.Next
		fast = fast.Next
	}

	slow.Next = slow.Next.Next

	return head
}

/*
	求链表的中间节点
*/
func  FindMiddleNode(head *LinkNode) *LinkNode {
	if nil == head || nil == head.Next {
		return nil
	}
	if nil == head.Next.Next {
		return head
	}
	
	// fast 一次走两步，slow一次走一步，那么fast走完全程的时候，slow应该是走到一半
	slow, fast := head, head
	for nil != fast && nil != fast.Next {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func printNode(head *LinkNode) {
	for head != nil {
		//fmt.Print(head.value, "\t")
		fmt.Println(head)
		head = head.Next
	}
	fmt.Println()
}

func getNodeList() *LinkNode {
	node1 := new(LinkNode)
    node1.Data = 1
    node2 := new(LinkNode)
    node2.Data = 2
    node3 := new(LinkNode)
    node3.Data = 3
    node4 := new(LinkNode)
	node4.Data = 4
	node5 := new(LinkNode)
	node5.Data = 5
	node6 := new(LinkNode)
    node6.Data = 6
    node1.Next = node2
    node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node6
	// node6.Next = node3
	
	return node1
}









