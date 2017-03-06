package main

import(
    "fmt"
)

func main(){
    slice_int := make([]int,0)
    slice_int = append(slice_int,1,2,3)
    slice_int = append(slice_int,4,5,6)
    fmt.Println(slice_int)

    slice_int2 := make([]int,0)
    slice_int2 = append(slice_int2,7,8,9)
    slice_int = append(slice_int,slice_int2...) //注意3个点
    fmt.Println(slice_int)

    slice_string := make([]string,0)
    slice_string = append(slice_string,"hello")
    slice_string = append(slice_string,"world")
    fmt.Println(slice_string)

}