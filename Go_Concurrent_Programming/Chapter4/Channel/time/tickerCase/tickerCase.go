/**
  * Author: Juntaran
  * Email:  Jacinthmail@gmail.com
  * Date:   2017/5/11 01:45
  */

package main

import (
	"time"
	"fmt"
)

// 断续器  var ticker *time.Ticker = time.NewTicker(time.Second)
// 断续器一旦初始化，所有的绝对到期时间就已经确定了，断续器适合定时任务触发

func main() {
	intChan := make(chan int, 1)
	ticker := time.NewTicker(time.Second)

	go func() {
		for _ = range ticker.C {
			select {
			case intChan <- 1:
			case intChan <- 2:
			case intChan <- 3:
			}
		}
		fmt.Println("End. [sender]")
	} ()

	var sum int
	for e := range intChan {
		fmt.Printf("Received: %v\n", e)
		sum += e
		if sum > 10 {
			fmt.Printf("Got: %v\n", sum)
			break
		}
	}
	fmt.Println("End. [receiver]")
}
