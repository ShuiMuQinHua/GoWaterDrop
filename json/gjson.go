package main

import (
	"github.com/tidwall/gjson"
)

const json = `{"data":[{"name":{"first":"Janet","last":"Prichard"},"age":47},{"name":{"first":"Janet2","last":"Prichard2"},"age":48}]}`
const json1 = `{"data":[[11,null,33],[44,55,66]]}`

func main() {
	result := gjson.Get(json1, "data")
	result.ForEach(func(key,value gjson.Result)bool{
		res := value.Array()
		println(res[0].String())
		println(res[1].String())
		println(res[2].String())
		return true
	})
}