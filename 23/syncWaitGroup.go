package main

import (
	"fmt"
	"sync"
)

// WaitGroup，它用于线程同步，WaitGroup等待一组线程集合完成，才会继续向下执行
// 主线程(goroutine)调用Add来设置等待的线程(goroutine)数量
// 然后每个线程(goroutine)运行，并在完成后调用Done
// 同时，Wait用来阻塞，直到所有线程(goroutine)完成才会向下执行。
// Add(-1) 和 Done() 效果一致

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(t int) {
			defer wg.Done()
			fmt.Println(t)
		}(i)
	}

	// 如果移除此句, 程序看不到执行效果: 因为主程序先于 goroutine 结束了.
	wg.Wait()
}
