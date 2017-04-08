package main

import (
	"sync"
	"runtime"
	"fmt"
	"sync/atomic"
)

// 原子函数能够利用很底层的加锁机制来同步访问整型变量和指针

/*
	如果两个或多个goroutine在没有互相同步的情况下，访问某个共享的资源，并试图同时读写该资源时
	就处于互相竞争的状态。竞争状态让并发程序变得复杂。
	对一个共享资源的读写操作必须是原子化，同一时刻只有一个goroutine能够对共享资源进行读写操作
*/

var (
	// counter是本例中每个goroutine都要增加值的变量，即共享资源
	counter 	int64
	wg			sync.WaitGroup
)

func addCounter(id int)  {
	defer wg.Done()

	for count:=0; count<2; count++ {
		// 安全地对counter+1
		atomic.AddInt64(&counter, 1)
		// 当前goroutine从线程退出，并放回队列
		runtime.Gosched()
		fmt.Println("id: ", id)
	}
	fmt.Println(id, "Finished")
}

func main() {
	//runtime.GOMAXPROCS(2)

	wg.Add(2)

	go addCounter(1)
	go addCounter(2)

	wg.Wait()
	fmt.Println("Final Counter: ", counter)

	// 最终结果为4
}
