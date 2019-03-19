package main

import "fmt"

// 规则一，当defer被声明时，其参数就会被实时解析
func main() {
	var i int = 1

	// 输出: result1 => 2 (而不是 4)
	defer fmt.Println("result1 =>", func() int {
		return i * 2
	}())

	// 注： 虽然 defer 在函数结束后促发，但是参数 i 在上面的defer 里面已经被 解析了
	i++

	// 输出: result => 4 （这里才输出4）
	defer fmt.Println("result2 =>", func() int {
		return i * 2
	}())
}
