package main

import "fmt"

type myCh struct {
	aaa string
}

type myTest struct {
	AA []*myCh
}

func main(){

	a := &myCh{
		aaa : "cyy",
	}
	b := make([]*myCh,0,1)
	b = append(b,a)

	aa := &myTest{
		AA : b,
	}
	str := fmt.Sprintf("%v",aa.AA[0])

	fmt.Println(str)
}