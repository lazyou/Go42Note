package main

import (
	"fmt"
)

type A struct {
	Face int
}
type Aa A // 自定义新类型Aa，没有基础类型A的方法

func (a A) f() {
	fmt.Println("hi ", a.Face)
}

func main() {
	var s A = A{Face: 9}
	s.f()

	// 新类型不会拥有原基础类型所附带的方法
	// 17\type_customize.go:22:4: sa.f undefined (type Aa has no field or method f)
	var sa Aa = Aa{Face: 9}
	sa.f()
}
