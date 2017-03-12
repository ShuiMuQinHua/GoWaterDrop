package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Response1 struct {
	Page   int
	Fruits []string
}

type Response2 struct {
	Page   int      `json:"page"` //当以一个json格式返回的时候，这个字段的key会设置为page
	Fruits []string `json:"fruits"`
}

//func Marshal(v interface{}) ([]byte, error) //将各种数据类型转化为json数据类型
//func Unmarshal (data []byte, v interface{})error  //将json格式的字符串解析成对应的接口类型
func main() {
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	//这里将切片和字典编码为JSON数组或对象
	slcD := []string{"apple我我我", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	reslD := &Response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	reslB, _ := json.Marshal(reslD)
	fmt.Println(string(reslB))

	// 你可以使用tag来自定义编码后JSON键的名称
	res2D := &Response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	//解码json数据为GO数据
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	var dat map[string]interface{}

	//解码过程，并检测相关可能存在的错误
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	//为了使用解码后map里面的数据
	num := dat["num"].(float64)
	fmt.Println(num)

	//访问嵌套的数据需要一些类型装换
	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	//我们还可以将json数据解码为自定义数据类型
	str := `{"page":1,"fruits":["apple","peach"]}`
	res := &Response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)

}
