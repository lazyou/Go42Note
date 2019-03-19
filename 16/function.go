package main

import (
	"fmt"
	"time"
)

// 定义函数类型 funcType
type funcType func(time.Time)

func main() {
	// 方式一：直接赋值给变量
	f := func(t time.Time) time.Time {
		return t
	}
	fmt.Println(f(time.Now()))

	// 方式二：定义函数类型 funcType 变量 timer
	var timer funcType = CurrentTime
	timer(time.Now())

	// 先把 CurrentTime 函数转为 funcType 类型，然后传入参数调用
	funcType(CurrentTime)(time.Now())
	// 这种处理方式在Go 中比较常见
}

func CurrentTime(start time.Time) {
	fmt.Println(start)
}
