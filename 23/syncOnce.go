package main

import (
	"fmt"
	"sync"
	"time"
)

var once sync.Once

// sync.Once.Do(f func()) 能保证 once 只执行一次,这个 sync.Once 块只会执行一次。
func main() {
	for i, v := range make([]string, 10) {
		once.Do(onces)
		fmt.Println("v:", v, "---i:", i)
	}

	for i := 0; i < 10; i++ {
		go func(i int) {
			// 这里将不会被执行， 出发上面的 once.Do(onces) 被移除
			once.Do(onced)
			fmt.Println(i)
		}(i)
	}

	time.Sleep(4000)
}

func onces() {
	fmt.Println("onces")
}

func onced() {
	fmt.Println("onced")
}
