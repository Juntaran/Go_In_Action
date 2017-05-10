/**
  * Author: Juntaran
  * Email:  Jacinthmail@gmail.com
  * Date:   2017/5/11 01:37
  */

package main

import (
	"time"
	"fmt"
)

// 定时器判断超时

func main() {
	intChan := make(chan int, 1)
	go func() {
		time.Sleep(time.Second)
		intChan <- 1
	} ()

	select {
	case e := <-intChan:
		fmt.Printf("Received: %v\n", e)
	case <-time.NewTimer(time.Millisecond * 500).C:
		fmt.Println("Timeout!")
	}
}
