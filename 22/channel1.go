package main

import (
	"fmt"
	"time"
)

func main() {
	// 不使用带缓冲区的channel
	c := make(chan int)
	go send(c)
	go recv(c)
	time.Sleep(3 * time.Second)
	close(c)
}

// 只能向 chan 里 send 数据
func send(c chan<- int) {
	for i := 0; i < 10; i++ {
		fmt.Println("send readey ", i)
		c <- i
		fmt.Println("send ", i)
	}
}

// 只能接收 channel 中的数据
func recv(c <-chan int) {
	for i := range c {
		fmt.Println("received ", i)
	}
}

/*
程序输出：
send readey  0
send  0
send readey  1
received  0
received  1
send  1
send readey  2
send  2
send readey  3
received  2
received  3
send  3
send readey  4
send  4
send readey  5
received  4
received  5
send  5
send readey  6
send  6
send readey  7
received  6
received  7
send  7
send readey  8
send  8
send readey  9
received  8
received  9
send  9
*/
