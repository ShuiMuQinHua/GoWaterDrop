package main

import(
    "bytes"
    "fmt"
    "encoding/json"
)


const json1 = `{
	"priority": 2,
	"hot_filter": 2,
	"deploy_copy_rate": 1,
	"deploy_copy_min": 5,
	"deploy_copy_max": 100,
	"deploy_ahead_rate": 0.5,
	"host_filter": [
		"haveTcpPort"
	]
}`

type myTest struct {
    Priority int `json:"priority"`
}

func main(){

    var t myTest
    err := json.Unmarshal([]byte(json1), &t)
    if err != nil {
        fmt.Println("err")
    }
    fmt.Println(t.Priority)
    return



    byte1 := []byte("mscdde") //字符串强制转换为byte切片
    byte2 := []byte("ms")
    spe1 := []byte("m") //申明一个byte切片 包含一个元素m
    spe2 := []byte("M")

    comres1 := bytes.Compare(spe2,spe1) //返回字典序比较的结果
    equres := bytes.Equal(spe2,spe1) //判断两个切片的内容是否完全相同
    equutf8res := bytes.EqualFold(spe2,spe1) //将unicode大写、小写、标题三种格式字符视为相同
    prefix := bytes.HasPrefix(byte1,byte2) //判断一个切片中是否有另一个切片前缀byte2
    suf := bytes.HasSuffix(byte1,byte2) //判断一个切片中是否有另一个切片后缀byte2
    contain := bytes.Contains(byte1,byte2) //包含
    
    fmt.Println(comres1)
    fmt.Println(equres)
    fmt.Println(equutf8res)
    fmt.Println(prefix)
    fmt.Println(suf)
    fmt.Println(contain)
    
    //type Reader
    //type Buffer
}