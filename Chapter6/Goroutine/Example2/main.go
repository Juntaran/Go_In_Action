package main

import (
	"sync"
	"runtime"
	"fmt"
)
// 这个例子消耗时间稍长，能够看出调度器切换两个goroutine

var wg sync.WaitGroup

// 显示5000以内的素数
func printPrime(number string)  {
	defer wg.Done()
next:
	for outer:=2; outer<5000; outer++ {
		for inner:=2; inner<outer; inner++ {
			if outer%inner == 0 {
				continue next
			}
		}
		fmt.Printf("%s:%d\n", number, outer)
	}
	fmt.Println("Completed", number)
}

func main() {
	// 分配一个逻辑处理器给调度器使用
	runtime.GOMAXPROCS(1)
	wg.Add(2)

	// 创建两个goroutine
	fmt.Println("Create Goroutines")
	go printPrime("A")
	go printPrime("B")

	// 等待goroutine结束
	fmt.Println("Waiting to Finish")
	wg.Wait()
	fmt.Println("Terminating Program")
}
