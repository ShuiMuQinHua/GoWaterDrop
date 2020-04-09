package main

import "fmt"

type user struct{
Name string
}

func main () {
	user := &user{
		Name:"aaaa",
	}
	
	test(user)
}


func test(u interface{}){
	if y, ok := u.(user);ok{
		fmt.Println(y.Name)
	}   //转成int，不报错
	
}
