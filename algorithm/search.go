package main

/* 查找算法 */

import (
	"fmt"
)

func main(){
	sli := getSortedSli()

	/* 二分查找 时间复杂度 O(logn) 
	flag := binarySearch(sli, 15, 0, len(sli)-1)
	fmt.Println(flag)
	*/

	/* 编程实现计算一个数的平方根 保留6位小数 */


	/* 快速定位IP地址对应的省份信息 转换为  查找最后一个小于等于给定值的元素 */

	/* 查找第一个值等于给定值的元素 对于有重复的元素的slice */
	index := binarySearchFirst(sli, 23)
	fmt.Println(index)

	/* 查找最后一个值等于给定值的元素 对于有重复的元素的slice */
	index = binarySearchLast(sli, 23)
	fmt.Println(index)

	/* 查找第一个大于等于给定值的元素 对于有重复的元素的slice */
	index = binarySearchMore(sli, 95)
	fmt.Println(index)

	/* 查找最后一个小于等于给定值的元素 对于有重复的元素的slice */
	index = binarySearchLess(sli, 56)
	fmt.Println(index)
}

/* 二分查找  时间复杂度 O(logn) */
func binarySearch(sli []int, v int, low int, high int) bool {
	if low > high {
		return false
	}

	if low == high && sli[low] != v {
		return false
	}

	m := low + ((high - low) >>1)

	if sli[m] == v {
		return true
	}
	
	if sli[m] > v {
		return binarySearch(sli, v, low, m-1)
	}

	if sli[m] < v {
		return binarySearch(sli, v, m+1, high)
	}

	return false
}

func binarySearchFirst(sli []int, v int) int {
	low := 0
	high := len(sli) - 1

	for ; low <= high ; {
		mid := low + ((high - low) >> 1)
		if sli[mid] > v {
			high = mid - 1
		}else if sli[mid] < v {
			low = mid + 1
		} else {
			if mid == 0 || sli[mid-1] != v {
				return mid
			} else {
				high = mid - 1
			}
		}
	}

	return -1
}

func binarySearchLast(sli []int, v int) int {
	low := 0
	high := len(sli) - 1

	for ; low <= high ; {
		mid := low + ((high - low) >> 1)
		if sli[mid] > v {
			high = mid - 1
		}else if sli[mid] < v {
			low = mid + 1
		} else {
			if mid == (len(sli) - 1) || sli[mid+1] != v {
				return mid
			} else {
				low = mid + 1
			}
		}
	}

	return -1
}

func binarySearchMore(sli []int, v int) int {
	low := 0
	high := len(sli) - 1

	for ;low <= high ; {
		mid := low + ((high - low) >> 1)
		if sli[mid] < v  {
			low = mid + 1
		} else   {
			if mid == 0 || sli[mid-1] < v {
				return mid
			} else  {
				high = mid - 1
			}
		} 
	}

	return -1
}

func binarySearchLess(sli []int, v int) int {
	low := 0
	high := len(sli) - 1

	for ;low <= high; {
		mid := low + ((high - low) >> 1)
		if sli[mid] > v {
			high = mid - 1
		} else {
			if mid == (len(sli) - 1) || sli[mid + 1] > v {
				return mid
			} else {
				low = mid + 1
			}
		}
	}

	return -1
}

func getSortedSli() []int {
	sli := []int{10,15,23,23,23,56,78,100,234,456,789,800,801,900}
	return sli
}

