package main

import (
	"fmt"
	"time"
)

type field struct {
	name string
}

func (p *field) print() {
	fmt.Println(p.name)
}

func main() {
	data := []field{{"one"}, {"two"}, {"three"}}

	// 如果 val 为指针，则会产生指针的拷贝，依旧可以修改集合中的原值
	// TODO: 可是我居然没看懂哪里改变值了, 不是只有打印嘛
	for _, v := range data {
		go v.print()
	}

	time.Sleep(3 * time.Second)
	// goroutines （可能）显示: three, three, three
}
