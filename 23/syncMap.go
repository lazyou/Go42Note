package main

import (
	"fmt"
	"sync"
)

// 随着Go1.9的发布，有了一个新的特性，那就是 sync.map，它是原生支持并发安全的 map。虽然说普通 map 并不是线程安全（或者说并发安全），但一般情况下我们还是使用它，因为这足够了；只有在涉及到线程安全，再考虑 sync.map。

// 但由于 sync.Map 的读取并不是类型安全的，所以我们在使用 Load 读取数据的时候我们需要做类型转换。

// sync.Map 的使用上和 map 有较大差异，详情见代码。
func main() {
	var m sync.Map

	// Store
	m.Store("name", "Joe")
	m.Store("gender", "Male")

	// LoadOrStore
	// 若key不存在，则存入key和value，返回false和输入的value
	v, ok := m.LoadOrStore("name1", "Jim")
	fmt.Println(ok, v) //false Jim

	// 若key已存在，则返回true和key对应的value，不会修改原来的value
	v, ok = m.LoadOrStore("name", "aaa")
	fmt.Println(ok, v) //true Joe

	// Load
	v, ok = m.Load("name")
	if ok {
		fmt.Println("key存在，值是： ", v)
	} else {
		fmt.Println("key不存在")
	}

	// Range
	// 遍历sync.Map
	f := func(k, v interface{}) bool {
		fmt.Println(k, v)
		return true
	}
	m.Range(f)

	// sDelete
	m.Delete("name1")
	fmt.Println(m.Load("name1"))
}
