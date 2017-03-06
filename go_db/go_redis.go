package main

import (
	"fmt"
	
	"github.com/garyburd/redigo/redis"
)

func main() {
	ope_redis()
}

func ope_redis() {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println(err)
		return
	}
	v, err := c.Do("SET", "name", "red")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(v)
	v, err = redis.String(c.Do("Get", "name"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(v)
}
