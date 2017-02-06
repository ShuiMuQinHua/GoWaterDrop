package main

import "fmt"

func main() {
	persion1 := map[string]string{"id": "12", "name": "wuyi", "age": "30"}

	persion2 := make(map[string]string)
	persion2["id"] = "13"
	persion2["name"] = "cyy"
	persion2["age"] = "30"

	fmt.Println(persion1)
	fmt.Println(persion2)
	fmt.Println(persion1["name"])
	fmt.Println(persion2["id"])

	/*循环map 使用key 输出map值*/
	for item := range persion1 {
		fmt.Println("persion1", item, "is", persion1[item])
	}

	/*查看元素在集合中是否存在*/
	itemvalue, ok := persion1["id"]
	/* 如果 ok 是 true, 则存在，否则不存在 */
	if ok {
		fmt.Println("persion1 id is", itemvalue)
	} else {
		fmt.Println("persion1 id is not exist")
	}

	/*删除元素值*/
	delete(persion1, "id")
	fmt.Println(persion1)
}
