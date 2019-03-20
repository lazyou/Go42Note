package main

import (
	"fmt"
)

// recover() 的调用仅当它在 defer 函数中被直接调用时才有效。
func div(a, b int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("捕获到错误：%s\n", r)
		}
	}()

	if b < 0 {

		panic("除数需要大于0")
	}

	fmt.Println("余数为：", a/b)

}

func main() {
	// 捕捉内部的Panic错误
	div(10, 0)

	// 捕捉主动Panic的错误
	div(10, -1)

	/*
		捕获到错误：runtime error: integer divide by zero
		捕获到错误：除数需要大于0
	*/
}
