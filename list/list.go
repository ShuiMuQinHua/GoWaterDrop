package main

import (
	"fmt"
	"container/list"
)

func main()  {
	li := list.New()
	li.PushFront("ccc")
	li.PushFront("bbb")
	li.PushFront("ddd")
	li.PushFront("eee")
	li.PushBack("fff")

	

	for e:= li.Front();e != nil;e = e.Next(){
		if e.Value == "bbb" {
			li.InsertBefore("mmm", e)
		}
		// fmt.Println( fmt.Sprintf("list element: %s", e.Value))
	}

	for e:= li.Front();e != nil;e = e.Next(){
		fmt.Println( fmt.Sprintf("list element: %s", e.Value))
	}
	 
	// li.Init()
	

	li.PushFrontList(li)
	li.PushBackList(li)
	fmt.Println( fmt.Sprintf("list len: %d", li.Len()))
}