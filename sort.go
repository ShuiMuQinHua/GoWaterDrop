package main

import(
    "fmt"
    "sort"
)

//使用冒泡排序 排列数据
func bubbleSort(data sort.IntSlice){
    // for pass := 1; pass < data.Len(); pass++ {
    //     for i := 0;i < data.Len() - pass; i++ {
    //         if data.Less(i+1, i) {
    //             data.Swap(i, i + 1)
    //         }
    //     }
    // }
    data.Sort()  //sort方法 就是使用上面的冒泡排序 实现的
    fmt.Println(data)
}

func main(){
    data := []int{2,5,3,6,3,7,9,8,100,56}
    bubbleSort(data)
}