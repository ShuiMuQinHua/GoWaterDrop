package main

import (
	"fmt"
	"strings"
)

func main() {
	tt := "ab,cT,EC,T"
	str1 := "2abc8242342"
	fmt.Println(strings.ToUpper(tt))
	fmt.Println(strings.Split(tt, ","))
	if strings.HasPrefix(str1, "abc") { //判断是否有前缀
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}

	if strings.HasSuffix(str1, "42") { //判断是否有后缀
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}

	if strings.Contains(str1, "c82") { //判断是否包含
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}

	ind1 := strings.Index(str1, "82") //匹配的字符串的第一个字符  所在的索引的位置  从0开始计数
	fmt.Println(ind1)

	ind2 := strings.LastIndex(str1, "4")
	fmt.Println(ind2) //返回最后一个匹配的位置的索引

	str2 := strings.Replace(str1, "2", "8", 2) //把str1中前两个2 变成8  并返回新的字符串
	fmt.Println(str2)

	ind3 := strings.Count(str1, "2") //计算在str1中  字符串“2”的出现次数
	fmt.Println(ind3)

	str3 := strings.Repeat(str1, 2) //把str1字符串重复两次  返回新的字符串
	fmt.Println(str3)

	str4 := strings.ToUpper(str1)
	fmt.Println(str4) //转换为大写

	str5 := strings.ToLower(str4)
	fmt.Println(str5) //转换为小写

	fmt.Println(str1)
	fmt.Println(strings.TrimSpace(str1))      //去除str1首尾的空白符号
	fmt.Println(strings.Trim(str1, "2"))      //去除str1中开头和结尾的 “2”字符串
	fmt.Println(strings.TrimLeft(str1, "2"))  //去除str1中开头的 "2"字符串
	fmt.Println(strings.TrimRight(str1, "2")) //去除str1中结尾的 "2"字符串
}
