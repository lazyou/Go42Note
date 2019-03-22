package main

import "fmt"

// 接口中内嵌结构体
type Writer interface {
	Write()
}

type Author struct {
	name string
	Writer
}

func (a Author) Write() {
	fmt.Println(a.name, "  Write.")
}

// 定义新结构体，重点是实现接口方法 Write()
type Other struct {
	i int
}

// 新结构体 Other 实现接口方法 Write()，也就可以初始化时赋值给 Writer 接口
func (o Other) Write() {
	fmt.Println(" Other Write.")
}

func main() {

	//  方法一：Other{99} 作为Writer 接口赋值
	Ao := Author{"Other", Other{99}}
	Ao.Write()

	// 方法二：简易做法，对接口使用零值，可以完成初始化
	Au := Author{name: "Hawking"}
	Au.Write()
}
