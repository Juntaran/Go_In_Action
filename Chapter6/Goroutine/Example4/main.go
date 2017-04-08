package main

import (
	"sync"
	"runtime"
	"fmt"
)

/*
	如果两个或多个goroutine在没有互相同步的情况下，访问某个共享的资源，并试图同时读写该资源时
	就处于互相竞争的状态。竞争状态让并发程序变得复杂。
	对一个共享资源的读写操作必须是原子化，同一时刻只有一个goroutine能够对共享资源进行读写操作
*/

var (
	// counter是本例中每个goroutine都要增加值的变量，即共享资源
	counter 	int
	wg			sync.WaitGroup
)

func addCounter(id int)  {
	defer wg.Done()

	for count:=0; count<2; count++ {
		// 捕获counter的值
		value := counter

		// 当前goroutine从线程退出，
		// runtime.Gosched()用于让出CPU时间片
		/*
			事实上，如果没有在代码中通过 runtime.GOMAXPROCS(n)
			其中n是整数，指定使用多核的话，
			goroutins都是在一个线程里的，
			它们之间通过不停的让出时间片轮流运行，达到类似同时运行的效果。
		*/
		runtime.Gosched()

		// 增加本地value变量
		value ++

		// value值保存给counter
		counter = value
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
}

// 因为两个goroutine互相竞争，最后counter本应为4却为2
// 为了解决这个问题，可以采用对共享资源加锁