/**
  * Author: Juntaran
  * Email:  Jacinthmail@gmail.com
  * Date:   2017/4/10 21:46
  */

package main

import (
	"sync"
	"math/rand"
	"time"
	"fmt"
)

const (
	numberGoroutines = 4		// 要使用的goroutine数量
	taskLoad         = 10		// 要处理的工作的数量
)

var wg sync.WaitGroup

func init() {
	// 初始化随机数种子
	rand.Seed(time.Now().Unix())
}

// worker处理从buffered channel传入的工作
func worker(tasks chan string, worker int)  {
	defer wg.Done()

	for {					// 死循环
		// 等待分配工作
		task, ok := <-tasks
		if !ok {
			// 有错误，此时通道是空的，而且已经关闭了
			fmt.Printf("Worker %d: Shutting Down\n", worker)
			return
		}
		// 开始工作
		fmt.Printf("Worker: %d Started %s\n", worker, task)

		// 随机等待一段时间模拟工作
		sleepTime := rand.Int63n(100)
		time.Sleep(time.Duration(sleepTime) * time.Millisecond * 100)

		// 工作已完成
		fmt.Printf("Worker: %d Completed %s\n", worker, task)
	}
}

func main() {
	// 创建一个有缓冲的通道来管理工作
	tasks := make(chan string, taskLoad)

	// 启动goroutine来处理工作
	wg.Add(numberGoroutines)

	for i:=1; i<=numberGoroutines; i++ {
		go worker(tasks, i)
	}

	// 增加一组要完成的工作
	for post:=1; post<=taskLoad; post++ {
		tasks <- fmt.Sprintf("Task: %d", post)
	}

	// 工作完成，退出通道
	close(tasks)		// 关闭通道后，goroutine仍然可以从通道接收数据，但是不能再向其中发送数据

	wg.Wait()
}
