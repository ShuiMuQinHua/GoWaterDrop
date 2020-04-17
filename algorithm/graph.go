package main

/* 图 
	1.顶点，边，度 (微信好友)
	2.有向图  入度，出度 (微博关注)
	4.带权图 (QQ好友亲密度)
	5.邻接矩阵存储方式 (稀疏图的时候利用效率就低)
	6.邻接表存储方式(很像散列表) (hash+链表的实现方式) 或者是 (hash+平衡查找二叉树) 后者更高效
*/

import (
	"fmt"
)

type Node struct{
	value int      //节点为int型
	searched bool
};

type Graph struct{
	nodes []*Node
	edges map[*Node][]*Node		//邻接表示的无向图
}


func main(){
	g := initGraph()
	// g.Print()


	/* 广度优先搜索算法 雷达扫描式  即先查找离起始顶点最近的，然后是次近的，依次往外搜索
	打印出从某一点出发的所有的边 判断是否有到达某一顶点的路径 
		g.BFS(g.nodes[0])
	*/
	
	/* 广度优先搜索算法 查找最短路径 
		g.BFS2(g.nodes[0], g.nodes[7])
		pre := make(map[*Node]*Node)
		g.BFS2DiGui(g.nodes[0], g.nodes[0], g.nodes[7], pre)
	*/

	/* 深度优先搜索算法 
	   最直观的例子，走迷宫
	*/
	// g.DFS()
	pre := make(map[*Node]*Node)
	g.DFS2DiGui(g.nodes[0], g.nodes[0], g.nodes[7], pre)


}

//增加节点
//可以理解为Graph的成员函数
func (g *Graph) AddNode(n *Node)  {
	g.nodes = append(g.nodes, n)
}

//增加边
func (g *Graph) AddEdge(u, v *Node) {
	g.edges[u] = append(g.edges[u],v)	//u->v边
	g.edges[v] = append(g.edges[v],u)	//u->v边
}

//打印图
func (g *Graph) Print(){
	//range遍历 g.nodes，返回索引和值
	for _,iNode:=range g.nodes{
		fmt.Printf("%v:",iNode.value)
		for _,next:=range g.edges[iNode]{
			fmt.Printf("%v->",next.value)
		}
		fmt.Printf("\n")
	}
}

func initGraph() Graph{
	g := Graph{}
	for i:=1;i<=9;i++{
		g.AddNode(&Node{i,false})
	}

	//生成边
	A := [...]int{1,1,2,2,2,3,4,5,5,6,1}
	B := [...]int{2,5,3,4,5,4,5,6,7,8,9}
	g.edges = make(map[*Node][]*Node)//初始化边
	for i:=0;i<11;i++{
		g.AddEdge(g.nodes[A[i]-1], g.nodes[B[i]-1])
	}
	return g
}

func (g *Graph) BFS(n *Node){
	var adNodes[] *Node		//存储待搜索节点
	n.searched = true
	fmt.Printf("%d:",n.value)
	for _,iNode:=range g.edges[n]{
		if !iNode.searched {
			adNodes = append(adNodes,iNode)
			iNode.searched=true
			fmt.Printf("%v ",iNode.value)
		}
	}
	fmt.Printf("\n")
	for _,iNode:=range adNodes{
		g.BFS(iNode)
	}
}


func (g *Graph) BFS2(n *Node, t *Node){
	if n == t {
		return
	}
	n.searched = true

	queue := make([]*Node, 0, 10)
	queue = append(queue, n)
	pre := make(map[*Node]*Node)

	j := 0
	for ;len(queue) != 0; {
		node := queue[j]
		for _,item:=range g.edges[node]{
			if !item.searched {
				pre[item] = node
				if item == t {
					myprint(pre, n, t)
					return
				}
				item.searched = true
				queue = append(queue, item)
			}
		}
		j++
	}
}

func (g *Graph) BFS2DiGui(s *Node, n *Node, t *Node, pre map[*Node]*Node){
	var adNodes []*Node		//存储待搜索节点
	n.searched = true
	
	for _,item:=range g.edges[n]{
		if !item.searched {
			pre[item] = n  //存的是 每一层的节点的上一层
			if item == t {
				myprint(pre, s, t)
				return
			}
			item.searched = true
			adNodes = append(adNodes, item)
		}
	}

	for _,iNode:=range adNodes{
		g.BFS2DiGui(s, iNode, t, pre)
	}
}

func myprint(pre map[*Node]*Node, s *Node, t *Node){
	exist,ok := pre[t]
	if ok && t != s {
		myprint(pre, s, exist)
	}
	fmt.Println(t)
}



func (g *Graph) DFS(){
	for _,iNode:=range g.nodes{
		if !iNode.searched{
			iNode.searched = true
			fmt.Printf("%v->",iNode.value)
			g.visitNode(iNode)
			fmt.Printf("\n")
			g.DFS()
		}
	}
}

func (g *Graph) DFS2DiGui(s *Node, n *Node, t *Node, pre map[*Node]*Node) {
	n.searched = true
	for _, item := range g.edges[n] {
		if !item.searched {
			pre[item] = n
			if item == t {
				myprint(pre,s,t)
				return
			}
			item.searched = true
			g.DFS2DiGui(s , item, t, pre)
		}
	}
}
 

func (g *Graph) visitNode(n *Node){
	for _,iNode:= range g.edges[n]{
		if !iNode.searched{
			iNode.searched = true
			fmt.Printf("%v->",iNode.value)
			g.visitNode(iNode)
			return
		}
	}
}

