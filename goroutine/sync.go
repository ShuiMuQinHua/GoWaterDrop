package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

func main() {
	max_concurrent_count := 20

	wg := sync.WaitGroup{}
	wg.Add(max_concurrent_count)

	for i := 0; i < max_concurrent_count; i++ {
		// 启动协程, 并发地发起 HTTP 请求.
		go func() {
			defer wg.Done()

			resp, _ := http.Get("http://www.baidu.com")

			defer resp.Body.Close()
			body, _ := ioutil.ReadAll(resp.Body)

			fmt.Println(string(body))
		}()
	}

	wg.Wait()
}

//func main() {
//	ch := make(chan string)

//	go sendData(ch)
//	go getData(ch)

//	time.Sleep(1e9)

//}

//func sendData(ch chan string) {
//	ch <- "www"
//	ch <- "ttt"
//	ch <- "lll"
//}

//func getData(ch chan string) {
//	var input string
//	//time.Sleep(1e9)
//	for {
//		input = <-ch
//		fmt.Printf("%s", input)
//	}
//}
