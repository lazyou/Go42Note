package main

import (
	"fmt"
	"time"
)

func main() {
	data := []string{"one", "two", "three"}

	// 当前的迭代变量作为匿名goroutine的参数。
	for _, v := range data {
		go func(str string) {
			fmt.Println(str)
		}(v)
	}

	time.Sleep(3 * time.Second)
	// goroutines输出: one, two, three
}
