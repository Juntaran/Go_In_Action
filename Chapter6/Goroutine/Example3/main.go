package main

import (
	"runtime"
	"sync"
	"fmt"
)

func main() {
	// 分配2个逻辑处理器给调度器使用 这个例子与Example1的区别就在于例1只分配了一个逻辑处理器
	runtime.GOMAXPROCS(2)
	//runtime.GOMAXPROCS(runtime.NumCPU())	// 给每个可用的核心分配一个逻辑处理器，NumCPU返回可以使用的物理处理器数

	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Start Groutines")

	// 声明一个匿名函数，并创建一个goroutine
	go func() {
		// 在函数退出的时候调用Done来通知main函数工作已完成
		defer wg.Done()

		// 显示字母表3次
		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	} ()

	// 声明一个匿名函数，并创建一个goroutine
	go func() {
		// 在函数退出的时候调用Done来通知main函数工作已完成
		defer wg.Done()

		// 显示字母表3次
		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	} ()

	// 等待goroutine结束
	fmt.Println("Waiting to Finish")
	// 一旦最后一个goroutine调用了Done，Wait方法会返回
	wg.Wait()
	fmt.Println("\nTerminating Program")
}