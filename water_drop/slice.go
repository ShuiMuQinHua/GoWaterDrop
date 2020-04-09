package main

import (
	"fmt"
	"strconv"
)

type phone struct {
	price  int
	color  string
	number int
}

func main() {
	a := 2
	b := "wjdiwj232"
	c := [3]string{"qweq", "wqrw", "dvnsdj"}  //数组
	d := []int{2, 4, 7, 9, 3}  //切片
	e := phone{price: 123, color: "red", number: 123456}  //结构体
	f := map[int]int{1: 2, 2: 5, 3: 7, 9: 8, 100: 68}  //map

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)

	var arr1 [6]int
	var slice1 []int = arr1[2:5] //  将arr1中从下标2到5-1 下的元素创建为一个新的切片

	fmt.Println("slice1 length " + strconv.Itoa(len(slice1)))  //int转string

	for i := 0; i < len(slice1); i++ {
		arr1[i] = i
	}
	
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}

	/*len() 和cap()函数练习*/
	var numbers = make([]int, 3, 5)   //初始化了一个切片  len=3 cap=5  slice=[0 0 0]

	printSlice(numbers)

	/*一个切片在未初始化之前默认为 nil，长度为 0*/
	var nilnumber []int
	printSlice(nilnumber)
	if nilnumber == nil {      //判断切片是空的的方法
		fmt.Println("切片是空的")
	}

	/*创建切片*/
	numslice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}  //创建并初始化了一个切片
	printSlice(numslice)

	/* 打印原始切片 */
	fmt.Println("numslice ==", numslice)

	/* 打印子切片从索引1(包含) 到索引4(不包含)*/
	fmt.Println("numslice[1:4] ==", numslice[1:4])  //将numslice中从下标1到4-1 下的元素创建为一个新的切片

	/* 默认下限为 0*/
	fmt.Println("numslice[:3] ==", numslice[:3]) //将numslice中从下标0到3-1 下的元素创建为一个新的切片

	/* 默认上限为 len(s)*/
	fmt.Println("numslice[4:] ==", numslice[4:])  //将numslice中从下标4到len-1 下的元素创建为一个新的切片

	numbers1 := make([]int, 0, 5)
	printSlice(numbers1)

	/* 打印子切片从索引  0(包含) 到索引 2(不包含) */
	number2 := numslice[:2]
	printSlice(number2)

	/* 打印子切片从索引 2(包含) 到索引 5(不包含) */
	number3 := numslice[2:5]
	printSlice(number3)

	/*	append() 和 copy() 函数*/
	appendcopy()
}

/*len() 和cap()函数练习*/
func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

/*	append() 和 copy() 函数*/
func appendcopy() {
	fmt.Println("切片copy与append")
	var numbers []int
	printSlice(numbers)

	/* 允许追加空切片 */
	numbers = append(numbers, 0)
	printSlice(numbers)

	/* 向切片添加一个元素 */
	numbers = append(numbers, 1)
	printSlice(numbers)

	/* 同时添加多个元素 */
	numbers = append(numbers, 2, 3, 4)
	printSlice(numbers)

	/* 创建切片 numbers1 是之前切片的两倍容量*/
	numbers1 := make([]int, len(numbers), (cap(numbers))*2)

	/* 拷贝 numbers 的内容到 numbers1 */
	copy(numbers1, numbers)
	printSlice(numbers1)
}
