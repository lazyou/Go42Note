package main

import "fmt"

func main() {
	switch1()
}

func switch1() {
	switch a := 10; {
	case a > 1:
		fmt.Println("a > 1")
		// 此处无需 break，虽然下一个 case 也满足条件，但并不会执行
	case a > 5:
		//case a == "字符串": // 每个 case 后面的类型都要与 a 一致
		fmt.Println("a > 5")
	default:
		fmt.Println("default")
	}
}
