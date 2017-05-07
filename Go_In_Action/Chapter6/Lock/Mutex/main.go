package main

import (
	"sync"
	"fmt"
	"runtime"
)

var (
	counter		int
	wg			sync.WaitGroup
	mutex		sync.Mutex			// mutex定义一段代码临界区
)

func addCounter(id int)  {
	defer wg.Done()

	for count:=0; count<2; count++ {
		// 同一时刻只允许一个goroutine进入这个拦截去
		mutex.Lock()
		{
			value := counter
			runtime.Gosched()		// 这里强制退出当前goroutine时，调度器会再次分配这个goroutine
			value ++
			counter = value
			fmt.Println("id: ", id)
		}
		mutex.Unlock()
		// 解锁，允许其他goroutine进入临界区
	}
}

func main() {
	wg.Add(2)

	go addCounter(1)
	go addCounter(2)

	wg.Wait()
	fmt.Printf("Final Counter: %d\n", counter)
}