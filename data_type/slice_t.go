package main

import (
	"fmt"
)

func main(){
	/* 1. map循环是随机的  2. 切片的 for range 会复制元素
		slice := []int{0, 1, 2, 3}
		m := make(map[int]*int)
		for key, val := range slice {
			m[key] = &val //因为这块儿赋值是给的指针，所以最后都会被覆盖为最后一个值3
		}
		for k, v := range m {
		fmt.Println(k, "->", *v)   //输出结果 v 都是3  k是 随机  0 1 2 3
		}

		n := make(map[int]int)
		for key,val := range slice {
			n[key] = val
		}
		for k, v := range n {
			fmt.Println(k, "->", v)
		}
		return
	*/
	
	/* 切片的长度和容量 
		slice1 := make([]int,0,5) //创建一个slice 长度为0, 容量为5
		slice1 = append(slice1,1)
		slice1 = append(slice1,2)
		slice1 = append(slice1,3)
		fmt.Println(slice1)  // 1 2 3

		s := make([]int, 5) //创建一个slice 长度是5
		s = append(s, 1, 2, 3)
		fmt.Println(s)  // 0 0 0 0 0 1 2 3
		
		s1 := make([]int, 0) //创建一个slice 长度是0
		s1 = append(s1, 1, 2, 3, 4)
		fmt.Println(s1) // 1 2 3 4
		return
	*/
	
	/* slice 底层append的时候，如果长度超过最大容量，会重新分配一块内存 
		sli1 := []int{1,2,3}
		sli2 := sli1
		
		//调换这两个语句的顺序，有很大的影响。 如果先append 会导致底层重新分配一块内存给sli2，这样后面再修改1的数据，就不再会影响sli1
		sli2[1] = 8
		sli2 = append(sli2, 8) 

		fmt.Println(sli1) // 1 8 3
		fmt.Println(sli2) // 1 8 3 8
		return
	*/

	/* slice 裁剪 和追加 
		arr := [5]int{1, 2, 3, 4, 5}
		s := arr[:]
		t := s[1:2]  // [ startIndex: endIndex] 包含startIndex 不包含endIndex
		fmt.Println(t)  // [2]

		fmt.Println(s[:0]) // []

		s1 := []int{6, 7, 8}
		s = append(s, s1...)
		fmt.Println(s)
		return
	*/


}