package main

import "fmt"

type Human struct {
	name string
}

type Student struct { // 含内嵌结构体Human
	Human // 匿名（内嵌）字段
	int   // 匿名（内嵌）字段
}

func main() {
	s := new(Student)
	s.int = 1
	fmt.Println(s)
}
