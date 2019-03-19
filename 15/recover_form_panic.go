package main

import (
	"log"
)

// Recover：从 panic 中恢复
func main() {
	protect(func() {
		panic("Recover：从 panic 中恢复")
	})
}

// protect() 函数接收一个匿名函数作为参数
func protect(g func()) {
	defer func() {
		log.Println("done")

		// 即使有panic，Println也正常执行。
		if err := recover(); err != nil {
			log.Printf("run time panic: %v", err)
		}
	}()

	log.Println("start")

	g() //   可能发生运行时错误的地方
}
