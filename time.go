package main

import (
	"fmt"
	"time"
)

const (
	f_date        = "2006-01-02"          //长日期格式
	f_shortdate   = "06-01-02"            //短日期格式
	f_times       = "15:04:02"            //长时间格式
	f_shorttime   = "15:04"               //短时间格式
	f_datetime    = "2006-01-02 15:04:02" //日期时间格式
	f_newdatetime = "2006/01/02 15~04~02" //非标准分隔符的日期时间格式
	f_newtime     = "15~04~02"            //非标准分隔符的时间格式
)

func main() {
	//我们要用上面的各种格式来格式化这个thisdate
	thisdate := "2014-03-17 14:55:06"
	timeformatdate, _ := time.Parse(f_datetime, thisdate)
	fmt.Println(timeformatdate)

	convdate := timeformatdate.Format(f_date)
	fmt.Println(convdate)

	convshortdate := timeformatdate.Format(f_shortdate)
	fmt.Println(convshortdate)

	convtime := timeformatdate.Format(f_times)
	fmt.Println(convtime)

	convshorttime := timeformatdate.Format(f_shorttime)
	fmt.Println(convshorttime)

	convnewdatetime := timeformatdate.Format(f_newdatetime)
	fmt.Println(convnewdatetime)

	convnewtime := timeformatdate.Format(f_newtime)
	fmt.Println(convnewtime)

	t := time.Now()                              //获取当前时间的结构体
	fmt.Println(t.Format("2006-01-02 15:04:05")) //格式化输出
	fmt.Println(t.Unix())                        //获取当前时间的时间戳

	fmt.Println(t.Day())    //获取当前时间的号
	fmt.Println(t.Minute()) //获取当前时间的分钟
	fmt.Println(t.Year())   //获取当前时间的年份
}
