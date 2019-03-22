package main

import "fmt"

func main() {
	type S struct {
		a int
		b float64
	}

	// new(S)为S类型的变量分配内存，并初始化（a = 0，b = 0.0），返回包含该位置地址的类型* S的值。
	sByNew := new(S)
	sByNew.a = 2
	fmt.Println(sByNew)

	// 给 s 分配内存，并零值化内存，但是这个时候 s 是类型 S
	var s S
	s.a = 1
	fmt.Println(s)
}
