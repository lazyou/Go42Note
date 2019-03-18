package main

import "fmt"

// 如果在执行完每个分支的代码后，还希望继续执行后续分支的代码，可以使用 fallthrough 关键字来达到目的
// 就是觉得没啥用, 还特别误导人。。。
func main() {
	switch a := 1; {
	case a == 1:
		fmt.Println("The integer was == 1")
		fallthrough
	case a == 2:
		fmt.Println("The integer was == 2")
	case a == 3:
		fmt.Println("The integer was == 3")
		fallthrough
	case a == 4:
		fmt.Println("The integer was == 4")
	case a == 5:
		fmt.Println("The integer was == 5")
		fallthrough
	default:
		fmt.Println("default case")
	}
}
