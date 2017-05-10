/**
  * Author: Juntaran
  * Email:  Jacinthmail@gmail.com
  * Date:   2017/5/11 01:30
  */

package main

import (
	"time"
	"fmt"
)

func main() {
	// 初始化一个2秒定时器
	timer := time.NewTimer(2 * time.Second)
	fmt.Printf("Present time: %v.\n", time.Now())
	// .C是一个带缓冲的通道
	expirationTime := <-timer.C
	fmt.Printf("Expiration time: %v.\n", expirationTime)
	fmt.Printf("Stop timer: %v.\n", timer.Stop())
}
