package main

import (
	"fmt"
	"container/heap"
)

type MyStudy struct {

}

func (m *MyStudy) Push(v interface{}) {

}

func (m *MyStudy) Pop() interface{} {
	return nil
}

func (m *MyStudy) Len() int {
	return 0
}

func (m *MyStudy) Less(i,j int) bool {
	return i>j
}

func (m *MyStudy) Swap(i,j int) {

}

