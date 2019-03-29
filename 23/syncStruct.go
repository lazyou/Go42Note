package main

import (
	"fmt"
	"sync"
	"time"
)

type Book struct {
	BookName string
	L        *sync.Mutex
}

func (bk *Book) SetName(wg *sync.WaitGroup, name string) {
	defer func() {
		fmt.Println("Unlock set name:", name)
		// 释放锁
		bk.L.Unlock()
		wg.Done()
	}()

	// 加锁
	bk.L.Lock()
	fmt.Println("Lock set name:", name)

	time.Sleep(1 * time.Second)
	bk.BookName = name
}

// Mutex 也可以作为 struct 的一部分，这样这个 struct 就会防止被多线程更改数据。
func main() {
	bk := Book{}
	bk.L = new(sync.Mutex)
	wg := &sync.WaitGroup{}
	books := []string{"《三国演义》", "《道德经》", "《西游记》"}

	for _, book := range books {
		wg.Add(1)
		go bk.SetName(wg, book)
	}

	wg.Wait()
}
