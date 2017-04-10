/**
  * Author: Juntaran
  * Email:  Jacinthmail@gmail.com
  * Date:   2017/4/10 16:53
  */

package main

import (
	"sync"
	"fmt"
	"time"
)

var wg sync.WaitGroup

// Runner模拟接力比赛的一个跑步者
func Runner(baton chan int)  {
	var newRunner int

	// 等待接力棒，runner接收baton通道的值
	runner := <- baton

	// 开始跑步
	fmt.Printf("Runner %d Running with Baton\n", runner)

	// 创建下一个跑步者
	if runner != 4 {
		newRunner = runner + 1
		fmt.Printf("Runner %d to the line\n", newRunner)
		go Runner(baton)
	}

	// 模拟跑步
	time.Sleep(1000 * time.Millisecond)

	// 判断比赛是否结束
	if runner == 4 {
		fmt.Printf("Runner %d Finished, Race Over\n", runner)
		wg.Done()
		return
	}

	// 比赛没结束的话接力棒交给下一个人
	fmt.Printf("Runner %d Exchange With Runner %d\n", runner, newRunner)

	baton <- newRunner
}


func main() {
	// 创建一个无缓冲的通道
	baton := make(chan int)

	// 为最后一个跑步者计数+1
	wg.Add(1)

	// 第一个跑步者持有接力棒
	go Runner(baton)

	// 开始比赛
	baton <- 1

	// 等待结束
	wg.Wait()
}
