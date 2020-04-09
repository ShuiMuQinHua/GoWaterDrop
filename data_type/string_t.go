package main

import (
	"fmt"
	"strconv" //处理string类型的转换
	"strings"
)

func main(){

	tt := "ab,cT,EC,T"
	str1 := "2ab c82  42342"
	str0 := "sakdsai,iwrewi,0976767"
	const a = `0.63519283478325730257
73547328325673456348767348` //Go中多行的连接

	//在Go当中 string底层是用byte数组存的，并且是不可以改变的 中文字符是用3个字节存的
	//要想访问中文的话，还是要用rune切片，这样就能按下标访问
	str000 := "go编程"
	c := "siueijfierjfierg" +
		"sjdusjfuesrjfue"
	fmt.Println(c)
	fmt.Println(a)
	fmt.Println(len(str000))         //8 因为一个汉字 3个字节
	fmt.Println(len([]rune(str000))) //4

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
	fmt.Println(strings.Fields(str1))         //把str1 以空格分隔  返回一个切片 slice
	fmt.Println(strings.Split(str0, ","))     //用, 去分隔str0 返回一个切片  slice

	slic := []string{"we32", "242dfd", "2343df"} //如何规定了长度就是数组，如果没规定长度  就是切片
	fmt.Println(strings.Join(slic, ","))         //用,把切片slic连接成字符串

	/*
		string的类型转换
	*/
	fmt.Println(strconv.IntSize) //用于获取程序所运行的平台下 int类型所占的位数
	in1 := 100876
	fmt.Println(strconv.Itoa(in1)) //数字转换为字符串
	//	strconv.FormatFloat()          //将64位浮点型的数据转换为字符串

	str00 := "23432"
	in11, _ := strconv.Atoi(str00) //字符串转换为数字
	fmt.Println(in11)

	flo64, _ := strconv.ParseFloat(str00, 64) //将字符串转换为float64
	fmt.Println(flo64)


	fmt.Println("查找子串是否在指定的字符串中")
	fmt.Println(" Contains 函数的用法")
	fmt.Println(strings.Contains("seafood", "foo")) //true
	fmt.Println(strings.Contains("seafood", "bar")) //false
	fmt.Println(strings.Contains("seafood", ""))    //true
	fmt.Println(strings.Contains("", ""))           //true 这里要特别注意
	fmt.Println(strings.Contains("我是中国人", "我"))     //true
	fmt.Println("")

	//判断字符串s是否包含字符串chars中的任一字符
	fmt.Println(" ContainsAny 函数的用法")
	fmt.Println(strings.ContainsAny("team", "i"))        // false
	fmt.Println(strings.ContainsAny("failure", "u & i")) // true
	fmt.Println(strings.ContainsAny("foo", ""))          // false
	fmt.Println(strings.ContainsAny("", ""))             // false
	fmt.Println("")
	fmt.Println(" ContainsRune 函数的用法")
	fmt.Println(strings.ContainsRune("我是中国", '我')) // true 注意第二个参数，用的是字符
	fmt.Println("")
	fmt.Println(" Count 函数的用法")
	fmt.Println(strings.Count("cheese", "e")) // 3
	fmt.Println(strings.Count("five", ""))    // before & after each rune result: 5 , 源码中有实现
	fmt.Println("")
	fmt.Println(" EqualFold 函数的用法")
	fmt.Println(strings.EqualFold("Go", "go")) //大小写忽略
	fmt.Println("")
	fmt.Println(" Fields 函数的用法")
	//["foo" "bar" "baz"] 返回一个列表
	fmt.Println("Fields are: %q", strings.Fields("  foo bar  baz   "))
	//相当于用函数做为参数，支持匿名函数
	for _, record := range []string{" aaa*1892*122", "aaa\taa\t", "124|939|22"} {
		fmt.Println(strings.FieldsFunc(record, func(ch rune) bool {
			switch {
			case ch > '5':
				return true
			}
			return false
		}))
	}
	fmt.Println("")
	fmt.Println(" HasPrefix 函数的用法")
	fmt.Println(strings.HasPrefix("NLT_abc", "NLT")) //前缀是以NLT开头的
	fmt.Println("")
	fmt.Println(" HasSuffix 函数的用法")
	fmt.Println(strings.HasSuffix("NLT_abc", "abc")) //后缀是以NLT开头的
	fmt.Println("")
	fmt.Println(" Index 函数的用法")
	fmt.Println(strings.Index("NLT_abc", "abc")) // 返回第一个匹配字符的位置，这里是4
	fmt.Println(strings.Index("NLT_abc", "aaa")) // 在存在返回 -1
	fmt.Println(strings.Index("我是中国人", "中"))     // 在存在返回 6
	fmt.Println("")
	fmt.Println(" IndexAny 函数的用法")
	fmt.Println(strings.IndexAny("我是中国人", "中")) // 在存在返回 6
	fmt.Println(strings.IndexAny("我是中国人", "和")) // 在存在返回 -1
	fmt.Println("")
	fmt.Println(" Index 函数的用法")
	fmt.Println(strings.IndexRune("NLT_abc", 'b')) // 返回第一个匹配字符的位置，这里是4
	fmt.Println(strings.IndexRune("NLT_abc", 's')) // 在存在返回 -1
	fmt.Println(strings.IndexRune("我是中国人", '中'))   // 在存在返回 6
	fmt.Println("")
	fmt.Println(" Join 函数的用法")
	s := []string{"foo", "bar", "baz"}
	fmt.Println(strings.Join(s, ", ")) // 返回字符串：foo, bar, baz
	fmt.Println("")
	fmt.Println(" LastIndex 函数的用法")
	fmt.Println(strings.LastIndex("go gopher", "go")) // 3
	fmt.Println("")
	fmt.Println(" LastIndexAny 函数的用法")
	fmt.Println(strings.LastIndexAny("go gopher", "go")) // 4
	fmt.Println(strings.LastIndexAny("我是中国人", "中"))      // 6
	fmt.Println("")
	fmt.Println(" Map 函数的用法")
	rot13 := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			return 'A' + (r-'A'+13)%26
		case r >= 'a' && r <= 'z':
			return 'a' + (r-'a'+13)%26
		}
		return r
	}
	fmt.Println(strings.Map(rot13, "'Twas brillig and the slithy gopher..."))
	fmt.Println("")
	fmt.Println(" Repeat 函数的用法")
	fmt.Println("ba" + strings.Repeat("na", 2)) //banana
	fmt.Println("")
	fmt.Println(" Replace 函数的用法")
	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1))
	fmt.Println("")
	fmt.Println(" Split 函数的用法")
	fmt.Printf("%q\n", strings.Split("a,b,c", ","))
	fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a "))
	fmt.Printf("%q\n", strings.Split(" xyz ", ""))
	fmt.Printf("%q\n", strings.Split("", "Bernardo O'Higgins"))
	fmt.Println("")
	fmt.Println(" SplitAfter 函数的用法")
	fmt.Printf("%q\n", strings.SplitAfter("/home/m_ta/src", "/")) //["/" "home/" "m_ta/" "src"]
	fmt.Println("")
	fmt.Println(" SplitAfterN 函数的用法")
	fmt.Printf("%q\n", strings.SplitAfterN("/home/m_ta/src", "/", 2))  //["/" "home/m_ta/src"]
	fmt.Printf("%q\n", strings.SplitAfterN("#home#m_ta#src", "#", -1)) //["/" "home/" "m_ta/" "src"]
	fmt.Println("")
	fmt.Println(" SplitN 函数的用法")
	fmt.Printf("%q\n", strings.SplitN("/home/m_ta/src", "/", 1))
	fmt.Printf("%q\n", strings.SplitN("/home/m_ta/src", "/", 2))  //["/" "home/" "m_ta/" "src"]
	fmt.Printf("%q\n", strings.SplitN("/home/m_ta/src", "/", -1)) //["" "home" "m_ta" "src"]
	fmt.Printf("%q\n", strings.SplitN("home,m_ta,src", ",", 2))   //["/" "home/" "m_ta/" "src"]
	fmt.Printf("%q\n", strings.SplitN("#home#m_ta#src", "#", -1)) //["/" "home/" "m_ta/" "src"]
	fmt.Println("")
	fmt.Println(" Title 函数的用法") //这个函数，还真不知道有什么用
	fmt.Println(strings.Title("her royal highness"))
	fmt.Println("")
	fmt.Println(" ToLower 函数的用法")
	fmt.Println(strings.ToLower("Gopher")) //gopher
	fmt.Println("")
	fmt.Println(" ToLowerSpecial 函数的用法")
	fmt.Println("")
	fmt.Println(" ToTitle 函数的用法")
	fmt.Println(strings.ToTitle("loud noises"))
	fmt.Println(strings.ToTitle("loud 中国"))
	fmt.Println("")
	fmt.Println(" Replace 函数的用法")
	fmt.Println(strings.Replace("ABAACEDF", "A", "a", 2)) // aBaACEDF
	//第四个参数小于0，表示所有的都替换， 可以看下golang的文档
	fmt.Println(strings.Replace("ABAACEDF", "A", "a", -1)) // aBaaCEDF
	fmt.Println("")
	fmt.Println(" ToUpper 函数的用法")
	fmt.Println(strings.ToUpper("Gopher")) //GOPHER
	fmt.Println("")
	fmt.Println(" Trim  函数的用法")
	fmt.Printf("[%q]", strings.Trim(" !!! Achtung !!! ", "! ")) // ["Achtung"]
	fmt.Println("")
	fmt.Println(" TrimLeft 函数的用法")
	fmt.Printf("[%q]", strings.TrimLeft(" !!! Achtung !!! ", "! ")) // ["Achtung !!! "]
	fmt.Println("")
	fmt.Println(" TrimSpace 函数的用法")
	fmt.Println(strings.TrimSpace(" \t\n a lone gopher \n\t\r\n")) // a lone gopher
}