package main

/* 各种排序算法的实现 */

import (
	"fmt"
)

func main(){
	sli := getToSortSlice()

	/* 冒泡排序 O(n*n)  稳定排序 原地排序
		sli = bubbleSort(sli)
		printSlice(sli)
	*/

	/* 插入排序  O(n*n) 稳定排序 原地排序
		sli = insertionSort(sli)
		printSlice(sli)
	*/

	/* 选择排序 O(n*n) 非稳定排序 原地排序
		sli = selectionSort(sli)
		printSlice(sli)
	*/
	
	/* 归并排序 O(n * logn) 稳定排序 非原地排序  空间复杂度O(n)
	sli = mergeSort(sli)
	printSlice(sli)
	*/

	/* 快速排序 O(n * logn) 非稳定排序  原地排序 空间复杂度O(1)
	sli = quickSort(sli, 0, len(sli) - 1)
	printSlice(sli)
	*/

	/* 快速排序实现查找数组中的第n大的元素 
		1.分区为  A[0..1..p]  A[p]  A[p...r]
		2.如果p==k 则，第n大的元素就是A[p]
		3.如果k>p  则在A[p...r]中找
		4.如果k<p  则再A[0..1..p]中找
	*/
	ret := findTopN(sli, 3-1, 0, len(sli) - 1)
	fmt.Println(ret)

	/* 桶排序(大日志文件排序)  O(n)*/
	bucketSort(sli)

	/* 计数排序(高考查分) O(n) */
	countSort(sli)

	/* 思考 GO的排序函数是如何实现的 */


}

/* 冒泡排序 
	1.时间复杂度(o(n*n)),空间复杂度是o(1)
	2.是稳定排序，值相等的元素，相对位置不会变更
	3.是一种原地排序算法，因为空间复杂度为o(1)
	4.一次冒泡，会归位一个元素
*/
func bubbleSort(sli []int) []int {
	for j:=0; j< len(sli);j++ {

		changeFlag := false

		//一次排序只能让一个元素归位
		for i:=0;i<(len(sli) - 1);i++ {
			if sli[i] > sli[i+1] {
				j := sli[i+1]
				sli[i+1] = sli[i]
				sli[i] = j
				changeFlag = true
			}
		}

		//如果某一次冒泡的过程中，没有发生元素的交换，就说明当前已经是有序了，不再需要接下来的循环冒泡
		if !changeFlag {
			break
		}

	}
	
	return sli
}

/* 插入排序
	1.时间复杂度(o(n*n)),空间复杂度是o(1)
	2.是稳定排序，值相等的元素，相对位置不会变更
	3.是一种原地排序算法，因为空间复杂度为o(1)
*/
func insertionSort(sli []int) []int {
	
	for i:=1; i<len(sli); i++ {
		valua := sli[i]
		j := i - 1

		for ; j>=0; j-- {
			if valua > sli[j] {
				sli[j+1] = sli[j]
			} else {
				break
			}
		}
		sli[j+1] = valua
	}
	return sli
}

/* 选择排序
	1.时间复杂度(o(n*n)),空间复杂度是o(1)
	2.不是稳定排序，值相等的元素，相对位置会变更
	3.是一种原地排序算法，因为空间复杂度为o(1)
*/
func selectionSort(sli []int) []int {
	length := len(sli)
    if length <= 1 {
        return sli
    }

    for i := 0; i < length; i++ {
        min := i // 初始的最小值位置从0开始，依次向右

        // 从i右侧的所有元素中找出当前最小值所在的下标
        for j := length - 1; j > i; j-- {
            if sli[j] < sli[min] {
                min = j
            }
        }
        //fmt.Printf("i:%d min:%d\n", i, min)

        // 把每次找出来的最小值与之前的最小值做交换
        sli[i], sli[min] = sli[min], sli[i]

        //fmt.Println(values)
	}
	
	return sli
}

/* 
	归并排序 时间复杂度 O(n * logn)
*/
func mergeSort(sli []int) []int  {
	if len(sli) <= 1 {
		return sli
	}
	//递[归]
	middle := len(sli) / 2
	//不断地进行左右对半划分
	left := mergeSort(sli[:middle])
	right := mergeSort(sli[middle:])
	//合[并]
	return merge(left, right)
}

func merge(left, right []int) (result []int) {
	l, r := 0, 0
    // 注意：[左右]对比，是指左的第一个元素，与右边的第一个元素进行对比，哪个小，就先放到结果的第一位，然后左或右取出了元素的那边的索引进行++
	for l < len(left) && r < len(right) {
        //从小到大排序.
		if left[l] > right[r] {
			result = append(result, right[r])
            //因为处理了右边的第r个元素，所以r的指针要向前移动一个单位
			r++
		} else {
			result = append(result, left[l])
            //因为处理了左边的第r个元素，所以r的指针要向前移动一个单位
			l++
		}
	}
    // 比较完后，还要分别将左，右的剩余的元素，追加到结果列的后面(不然就漏咯）。
	result = append(result, left[l:]...)
	result = append(result, right[r:]...)
	return
}


/*
	快排  时间复杂度 O(n * logn)
*/
func quickSort(a []int, lo, hi int) []int {
	if lo >= hi {
		return a
	}
	p := partition(a, lo, hi) //找到随机元素应该放在哪个位置
	quickSort(a, lo, p-1)  //对随机元素之前的数组排序
	quickSort(a, p+1, hi)  //对随机元素之后的数组排序

	return a
}


/*桶排序 
	大文件排序的时候可以用到，很好的思路。  每个桶的顺序已知，只需要每个桶内做快速排序，然后直接合并即可
	比如需要给一个10G的订单日志排序，而我们的机器就只有1G的内存，显然不能直接把所有的数据都
	读入内存做排序。  我们可以先扫描数据，分到n个文件中，这个n个文件的顺序是已知的，那么
	我们只需要针对每个文件做排序即可
	时间复杂度接近 O(n)  所以叫线性排序
*/
func bucketSort(sli []int) []int {
	return nil
}

/* 计数排序
	高考查分, 满分900，最小0分，我们把考生根据分数，分配到901个桶中，
	那每个桶内的分数就是相同的，不再需要排序
	时间复杂度 O(n) 所以叫线性排序
*/
func countSort(sli []int) []int {
	return nil
}


func partition(a []int, lo, hi int) int {
	pivot := a[hi] //选取最后一个元素作为分区点，分区点可以选择其他的  可以用三数取中法也可以用随机法来选择分区点
	i := lo - 1  //i指向的是pivot应该待的位置  
	for j := lo; j < hi; j++ {
		if a[j] < pivot {   //每次把比pivot小的元素前移交换
			//最终可以保证的是  小标<=i的元素，值都比pivot要小，也就是最终找到了pivot的位置
			//i指针肯定是后置于j指针的
			i++   
			a[j], a[i] = a[i], a[j]
		}
	}
	a[i+1], a[hi] = a[hi], a[i+1]
	return i + 1
}

/* 查找第k大的元素 */
func findTopN(a []int, k int, lo int, hi int) int {
	p := partition(a, lo, hi)
	
	if p == k-1 {
		return a[p]
	}

	if p < k {
		return findTopN(a, k-p+1, p+1, hi )
	}

	if p > k {
		return findTopN(a, k, 0 , p-1)
	}

	return 0
}


func getToSortSlice() []int {
	s := []int{87,34,20,80,45,67,29,36,49,46,3,100,233,46,36,34,86}
	return s
}

func printSlice(sli []int) {
	for i:=0; i<len(sli); i++ {
		fmt.Println(sli[i])
	}
	return
}
