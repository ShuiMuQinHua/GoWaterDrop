/* 
1. 递归实现 斐波那契数列
2. 递归实现 排课 每个课程的学习都有一个前后顺序
3. 递归实现 走楼梯，总共的走法数
递归问题的核心：
	1. 一个问题的解可以分解为几个子问题的解
	2. 这个问题与分解之后的子问题，除了数据规模不同，求解思路完全一样
	3. 存在递归终止条件
写递归代码的关键就是找到如何将大问题分解为小问题的规律，并且基于此写出递推公式，然后再推敲终止条件，最后将递推公式和终止条件翻译成代码
编写递归代码的关键是，只要遇到递归，我们就把它抽象成一个递推公式，不用想一层层的调用关系，不要试图用人脑去分解递归的每个步骤
*/

package main

import (
	"fmt"
	// "time"
	"sort"
)

const LIM = 41

var fibs [LIM]uint64
var resData map[int]int

/*
利用一个全局的数组，做了一个斐波那契数列的缓存，提高性能
*/
func main() {
	
	/*斐波那契数列
		var result uint64 = 0
		start := time.Now()
		for i := 0; i < LIM; i++ {
			result = fibonacci(i)
			fmt.Printf("fibonacci(%d) is: %d\n", i, result)
		}
		end := time.Now()
		delta := end.Sub(start)
		fmt.Printf("longCalculation took this amount of time: %s\n", delta)
	*/

	/*递归排课
		directedGraph()
	*/

	/*递归计算爬楼梯的方式 
		resData = make(map[int]int)
		ret := stairsNum(7)
		fmt.Println(ret)
		ret = stairsNumFor(7)
		fmt.Println(ret)
	*/

	/* 最终推荐人 */
	ret := finalReferrer()
	fmt.Println(ret)
}

func fibonacci(n int) (res uint64) {
	// memoization: check if fibonacci(n) is already known in array:
	if fibs[n] != 0 {
		res = fibs[n]
		return
	}
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	fibs[n] = res
	return
}

/* 递归实现排课程 每个课程的学习会有一个前后顺序 */
func directedGraph(){
	var prereqs = map[string][]string{
		"algorithms": {"data structures"},
		"calculus": {"linear algebra"},
		"compilers": {
			"data structures",
			"formal languages",
			"computer organization",
		},
		"data structures":       {"discrete math"},
		"databases":             {"data structures"},
		"discrete math":         {"intro to programming"},
		"formal languages":      {"discrete math"},
		"networks":              {"operating systems"},
		"operating systems":     {"data structures", "computer organization"},
		"programming languages": {"data structures", "computer organization"},
	}

    for i, course := range topoSort(prereqs) {
        fmt.Printf("%d:\t%s\n", i+1, course)
    }
}

/* 1.终止条件，大的map中，不存在的key即可返回
2.每一行都是需要问其他行，我的每一个元素存不存在 */
func topoSort(m map[string][]string) []string {
    var order []string
    seen := make(map[string]bool)
	var visitAll func(items []string)
	
    visitAll = func(items []string) {
        for _, item := range items {
            if !seen[item] {
                seen[item] = true
                visitAll(m[item])
                order = append(order, item)
            }
        }
    }
    var keys []string
    for key := range m {
        keys = append(keys, key)
	}
    sort.Strings(keys)
    visitAll(keys)
    return order
}


//递归方式的楼梯问题 f(n) = f(n-1) + f(n-2)  f(1)=1 f(2)=2
func stairsNum(n int) int {
	var res int
	if resData[n] != 0 {
		res = resData[n]
		return res
	}
	
	if n == 1 {
		res = 1
		resData[n] = res
		return res
	}

	if n == 2 {
		res = 2
		resData[n] = res
		return res
	}
	
	res = stairsNum(n-2) + stairsNum(n-1)
	resData[n] = res

	return res
}

//非递归方式的楼梯问题
func stairsNumFor(n int) int {
	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	ret := 0
	pre := 2
	prepare := 1
	for i:=3; i<=n; i++ {
		ret = pre + prepare
		prepare = pre
		pre = ret
	}

	return ret
}

// 获取最终推荐人
func finalReferrer() string {
	var relation = map[string]string{
		"b":"a",
		"c":"b",
		"d":"c",
	}

	ret := ""

	var find func (str string)
	find = func (str string) {
		if relation[str] == "" {
			ret = str
			return 
		}

		find(relation[str])
	}
	find("d")
	return ret
}




