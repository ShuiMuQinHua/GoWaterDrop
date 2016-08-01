package main

import "fmt"

func main() {
	a := map[string]string{"id": "12", "name": "wuyi", "age": "30"}
	b := make(map[string]string)
	b["id"] = "13"
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(a["name"])
	fmt.Println(b["id"])
}
