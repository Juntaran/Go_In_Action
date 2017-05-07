package main

import (
	"sync"
	"fmt"
	"time"
	"sync/atomic"
)

var (
	// shutdown是通知正在执行的goroutine停止工作的标志
	shutdown 	int64
	wg			sync.WaitGroup
)

func doWork(name string)  {
	defer wg.Done()
	for {
		fmt.Printf("Doing %s Work\n", name)
		time.Sleep(1000 * time.Millisecond)

		// 判断是否停止工作
		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Printf("Shutting %s Down\n", name)
			break
		}
	}
}

func main() {
	wg.Add(2)

	go doWork("A")
	go doWork("B")

	// goroutine执行时间
	time.Sleep(5 * time.Second)

	// 停止工作，设置shutdown标志
	fmt.Println("Shutdown Now")
	atomic.StoreInt64(&shutdown, 1)			// atomic.StoreInt 安全的写  atomic.LoadInt 安全的读

	wg.Wait()
}