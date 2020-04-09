package main

import (
    "log"
    "net/http"
    "os"
	"math/rand"
	"time"
    "github.com/chenjiandongx/go-echarts/charts"
)

var nameItems = []string{"衬衫", "牛仔裤", "运动裤", "袜子", "冲锋衣", "羊毛衫"}
var seed = rand.NewSource(time.Now().UnixNano())

func randInt() []int {
    cnt := len(nameItems)
    r := make([]int, 0)
    for i := 0; i < cnt; i++ {
        r = append(r, int(seed.Int63()) % 50)
    }
    return r
}

func handler(w http.ResponseWriter, _ *http.Request) {
    bar := charts.NewBar()
    bar.SetGlobalOptions(charts.TitleOpts{Title: "Bar-示例图"}, charts.ToolboxOpts{Show: true})
    bar.AddXAxis(nameItems).
        AddYAxis("商家A", randInt()).
        AddYAxis("商家B", randInt())
    f, err := os.Create("bar.html")
    if err != nil {
        log.Println(err)
    }
    bar.Render(w, f) // Render 可接收多个 io.Writer 接口
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/map", handlerMap)
    http.ListenAndServe(":8081", nil)
}



var mapData = map[string]float32{
    "北京":   float32(rand.Intn(150)),
    "上海":   float32(rand.Intn(150)),
    "深圳":   float32(rand.Intn(150)),
    "辽宁":   float32(rand.Intn(150)),
    "青岛":   float32(rand.Intn(150)),
    "山西":   float32(rand.Intn(150)),
    "陕西":   float32(rand.Intn(150)),
    "乌鲁木齐": float32(rand.Intn(150)),
    "齐齐哈尔": float32(rand.Intn(150)),
}




func handlerMap(w http.ResponseWriter, _ *http.Request) {
    mc := charts.NewMap("china")
    mc.SetGlobalOptions(charts.TitleOpts{Title: "Map-示例图"})
mc.Add("map", mapData)

f, err := os.Create("map.html")
    if err != nil {
        log.Println(err)
    }
    mc.Render(w, f) 
}
