package main

import (
	"fmt"
	// "encoding/json"
)

type data struct {  
	Name string
	
}

type testMap struct {
	data string
	dataMap map[string]interface{}
}

func main() {
	/* 以下两种方式都会修改dm中的字段，因为map是引用类型 
		dm := make(map[string]interface{})
		dm["ddd"] = 111
		dm["vvv"] = 222
		t := &testMap{
			data : "test",
			dataMap:dm,
		}
		updateDm(t.dataMap)
		fmt.Println(t)

		dm["ddd"] = 111
		y := testMap{
			data : "test",
			dataMap:dm,
		}
		updateDm(y.dataMap)
		fmt.Println(y)

		updataData(y.data) //这种不会直接改y的data数据，因为string是值类型
		fmt.Println(y)

		updateDataPoint(&y.data) //这种会直接改y的data数据，因为传的是string的指针
		fmt.Println(y)
	*/	
	
	/* 因为map元素无法取值，所以不能直接给结构体字段赋值, 解决办法是把map的元素类型改成指向结构体的指针 
		m := map[string]data {"x":{"one"}}
		// m["x"].Name = "two" //error   cannot assign to struct field m["x"].name in map
		myd := m["x"]
		myd.Name = "two"
		fmt.Println(myd)

		mm := map[string]*data {"X":{"one"}}
		mm["X"].Name = "what?" //???
		strB, _ :=  json.Marshal(mm)
		fmt.Println(string(strB))
		return
	*/
	
	/* map数据make之后，就不为nil, 取一个不存在的key,返回的会是元素类型的0值 
		mData := make(map[string]int) 
		if mData == nil {
			fmt.Println("nil")
			return
		}
		val := mData["cc"]
		fmt.Println(val)  //int类型的0值是0，所以这块儿的返回值是0

		var aa map[string]int
		if aa == nil {
			fmt.Println("nil")  // 没有make的map数据为nil
			return
		}
		if _,ok := aa["ss"];ok {
			fmt.Println("aaaa")
			return
		}

		return
	*/
	
	/* map中删除一个不存在的key不会报错，访问一个不存在的key，会返回元素类型的0值 
		s := make(map[string]int)
		delete(s, "h")
		fmt.Println(s["h"])   //输出 0
	*/
	
	/* map需要深度复制
		acc := make(map[string]int)
		acc["cyy"] = 111
		bcc := make(map[string]int)
		bcc = acc
		bcc["cyy"] = 222
		fmt.Println(acc)
		fmt.Println(bcc)
		return
	*/
	
}


func updateDm(dm map[string]interface{}) {
	dm["ddd"] = 888888
}

func updataData(str string) {
	str = "bbbbbbb"
}

func updateDataPoint(str *string) {
	*str = "aaaaa"
}
