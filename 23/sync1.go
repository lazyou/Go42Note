package main

import (
	"fmt"
	"sync"
	"time"
)

// 通过3个 goroutine 来体现 sync.Mutex 对资源的访问控制特征
func main() {
	wg := sync.WaitGroup{}
	var mutex sync.Mutex

	fmt.Println("Locking  (G0)")
	mutex.Lock()

	fmt.Println("locked (G0)")
	wg.Add(3)

	for i := 1; i < 4; i++ {
		go func(i int) {
			fmt.Printf("Locking (G%d)\n", i)
			mutex.Lock()
			fmt.Printf("locked (G%d)\n", i)

			time.Sleep(time.Second * 2)
			mutex.Unlock()
			fmt.Printf("unlocked (G%d)\n", i)
			wg.Done()
		}(i)
	}

	time.Sleep(time.Second * 5)
	fmt.Println("ready unlock (G0)")
	mutex.Unlock()
	fmt.Println("unlocked (G0)")
	wg.Wait()
}

// 通过程序执行结果我们可以看到，当有锁释放时，才能进行lock动作，G0锁释放时，才有后续锁释放的可能，这里是G1抢到释放机会。
/*
程序输出：
Locking  (G0)
locked (G0)
Locking (G1)
Locking (G3)
Locking (G2)
ready unlock (G0)
unlocked (G0)
locked (G1)
unlocked (G1)
locked (G3)
locked (G2)
unlocked (G3)
unlocked (G2)
*/
