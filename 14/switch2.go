package main

import "fmt"

func main() {
	switch2()
}

func switch2() {
	a := 1
	b := "string"

	// 此处每个 case 的类型不受限制, 因为 switch 后面没有任何表达式
	switch {
	case a == 1:
		fmt.Println("a == 1")
	case b == "string":
		fmt.Println("string")
	default:
		fmt.Println("default")
	}
}
