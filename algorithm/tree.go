package main

/* 树的相关联系 */


import (
	"fmt"
)

type TreeNode struct {
    Data int
	Left *TreeNode
	Right *TreeNode
}

func main(){
	tree := getTree()
	/* 二叉树
		1.满二叉树 所有层都达到最大节点数
		2.完全二叉树(除了最低层，其他层必须满，且最低层的叶子节点靠左) 适合用数组来存  因为可以直接根据父节点下标推算出左右子节点的下标  父i  左 2i 右 2i+1
	*/

	/* 二叉树的遍历  
		1.前序遍历  对于任意节点，先打印这个节点，再打印左子树，最后打印右子树
		2.中序遍历  对于任意节点，先打印左子树，再打印这个节点，最后打印右子树
		3.后续遍历  对于任意节点，先打印左子树，再打印右子树，最后打印这个节点
		查找、插入、删除等很多操作的时间复杂度都跟树的高度成正比，两个极端情况的时间复杂度分别是 O(n) 和 O(logn)，
		分别对应二叉树退化成链表的情况和完全二叉树。
	
		preOrder(tree)
		inOrder(tree)
		postOrder(tree)
	*/

	/* 二叉查找树 要求：在树中的任意一个节点，其左子树中的每个节点的值，都要小于这个节点的值，而右子树节点的值都大于这个节点的值 */
	flag := searchTree(tree, 7)
	fmt.Println(flag)

	/*  平衡二叉查找树 要求：二叉树中任意一个节点的左右子树的高度相差不能大于1(不一定)
		设计的目的是解决普通二叉查找树在频繁的插入、删除等动态更新的情况下，出现时间复杂度退化的问题，最终可能退化成一个链表
		平衡二叉查找树中“平衡”的意思，其实就是让整棵树左右看起来比较“对称”、比较“平衡”，不要出现左子树很高、右子树很矮的情况。
		这样就能让整棵树的高度相对来说低一些，相应的插入、删除、查找等操作的效率高一些
		1.红黑树
	*/


	/* 递归树 */
}

/* 前序遍历 */
func preOrder(node *TreeNode){
	if node == nil {
		return
	}
	fmt.Println(node.Data)
	preOrder(node.Left)
	preOrder(node.Right)
}

/* 中序遍历 */
func inOrder(node *TreeNode){
	if node == nil {
		return
	}
	preOrder(node.Left)
	fmt.Println(node.Data)
	preOrder(node.Right)
}

/* 后序遍历 */
func postOrder(node *TreeNode){
	if node == nil {
		return
	}
	preOrder(node.Left)
	preOrder(node.Right)
	fmt.Println(node.Data)
}

/* 二叉查找树查找 */
func searchTree(node *TreeNode, value int) bool  {
	if node == nil {
		return false
	}

	if node.Data == value {
		return true
	}

	if node.Data > value {
		return searchTree(node.Left, value)
	}

	if node.Data < value {
		return searchTree(node.Right, value)
	}

	return false
}

/* 二叉查找树插入 */
func searchTreeInsert(node *TreeNode, value int) bool {
	if node.Data > value && node.Left == nil {
		node.Left = &TreeNode{Data: value}
	}

	if node.Data < value && node.Right == nil {
		node.Right = &TreeNode{Data: value}
	}

	if node.Data > value {
		return searchTreeInsert(node.Left, value)
	}

	if node.Data < value {
		return searchTreeInsert(node.Right, value)
	}

	return false
}


/* 二叉查找树删除 */
func searchTreeDel(node *TreeNode, value int) bool {
	p := node  //待删除节点，初始化为根节点
	pp := &TreeNode{} //pp指向待删除节点的父节点

	for ; p != nil && p.Data != value ; {
		pp = p
		if p.Data > value {
			p = p.Left
		} else {
			p = p.Right
		}
	}
	if p == nil {
		return false  //没找到待删除的点
	}


	//删除节点是叶子节点，只有一个子节点
	if p.Left != nil {
		if pp.Data > p.Data {
			pp.Left = p.Left
		} else {
			pp.Right = p.Left
		}
	} else if p.Right != nil {
		if pp.Data > p.Data {
			pp.Left = p.Right
		} else {
			pp.Right = p.Right
		}
	} else {
		if pp.Data > p.Data {
			pp.Left = nil
		} else {
			pp.Right = nil
		}
	}

	//删除节点有2个子节点

	return false
}



func getTree() *TreeNode {
	node1 := &TreeNode{Data:1}
	node2 := &TreeNode{Data:2}
	node3 := &TreeNode{Data:3}
	node4 := &TreeNode{Data:4}
	node5 := &TreeNode{Data:5}
	node6 := &TreeNode{Data:6}
	node7 := &TreeNode{Data:7}
	node8 := &TreeNode{Data:8}
	node9 := &TreeNode{Data:9}
	node10 := &TreeNode{Data:10}

	// node1.Left = node2
	// node1.Right = node3
	// node2.Left = node4
	// node3.Right = node5
	// node4.Left = node6
	// node4.Right = node7
	// node5.Left  = node8
	// node6.Right = node9
	// node8.Left = node10

	node5.Left = node3
	node5.Right = node7

	node3.Left = node2
	node2.Left = node1
	node3.Right = node4

	node7.Left = node6
	node7.Right = node8
	node8.Right = node10
	node10.Left = node9

	return node5
	

	return node1

}
