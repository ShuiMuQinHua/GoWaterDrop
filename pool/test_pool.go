package main

import (
	"fmt"
	"sync"
)

type deployTest struct {
	Name  string
	Age int64
}

var respPool = sync.Pool{
	New: func() interface{} {
		return &deployTest{}
	},
}


func main(){
	dt := respPool.Get().(*deployTest)
	poolTest(dt)
	fmt.Println(dt)
	dt.Reset()
	respPool.Put(dt)
}

func poolTest(dt *deployTest) {
	dt.Name = "cyy"
	dt.Age = 23
	
	dt,err := getPool(dt)
	if err != nil {
		fmt.Println(err)
	}
	// dt.Reset()
	// respPool.Put(dt)

	
}

func getPool (dt *deployTest) (*deployTest,error) {
	dt.Name = "xxx"
	return dt,nil
}

func (dt *deployTest) Reset() {
	*dt = deployTest{}
}



