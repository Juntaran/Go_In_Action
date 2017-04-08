package main

import (
	"runtime"
	"sync"
	"fmt"
)

func main() {
	// 分配一个逻辑处理器给调度器使用
	runtime.GOMAXPROCS(1)

	// WaitGroup是一个计数信号量，用来记录运行的goroutine
	// WaitGroup大于0时，wait方法阻塞
	// wg用来等待程序完成，计数加2代表等待两个goroutine
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

// 第一个goroutine完成的速度太快了，以至于在调度器切换到第二个goroutine之前就完成了任务，所以先输出了所有大写字母，再输出小写字母