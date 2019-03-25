package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var quit chan int = make(chan int)

func loop() {
	for i := 0; i < 1000; i++ {
		Factorial(uint64(1000))
	}
	quit <- 1
}

func Factorial(n uint64) (result uint64) {
	if n > 0 {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}

var wg1, wg2 sync.WaitGroup

// 体验下Go语言中通过设定 runtime.GOMAXPROCS(2) ，也即手动指定CPU运行的核数，来体验多核CPU在并发处理时的威力。
// 不得不提，递归函数的计算很费CPU和内存，运行时可以根据电脑配置修改循环或递归数量。
func main() {
	fmt.Println("1:", time.Now())
	fmt.Println(runtime.NumCPU()) // 默认CPU核数
	a := 5000
	for i := 1; i <= a; i++ {
		wg1.Add(1)
		go loop()
	}

	for i := 0; i < a; i++ {
		select {
		case <-quit:
			wg1.Done()
		}
	}
	fmt.Println("2:", time.Now())
	wg1.Wait()
	fmt.Println("3:", time.Now())

	runtime.GOMAXPROCS(2) // 设置执行使用的核数
	a = 5000
	for i := 1; i <= a; i++ {
		wg2.Add(1)
		go loop()
	}

	for i := 0; i < a; i++ {
		select {
		case <-quit:
			wg2.Done()
		}
	}

	fmt.Println("4:", time.Now())
	wg2.Wait()
	fmt.Println("5:", time.Now())
}
