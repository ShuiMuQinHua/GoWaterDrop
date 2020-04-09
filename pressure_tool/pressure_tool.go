package main

import (
	"fmt"
	"flag"
)

var (
	clientNum  = flag.Int("c", 1000, "the number of concurrent")
	intervalTime  = flag.Int("d", 1, "the time interval")
	url = flag.String("d", "", "the url to request")
)

func main(){
	fmt.Println("begin start my pressure test tool !")

}