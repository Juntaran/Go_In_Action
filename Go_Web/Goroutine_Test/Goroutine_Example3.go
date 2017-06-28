package main

import (
	"time"
	"fmt"
)

func main() {
	exit := make(chan struct{})     // 创建通道
	go func() {
		time.Sleep(time.Second)
		fmt.Println("goroutine done.")
		close(exit)                 // 关闭通道，发出信号
	} ()

	fmt.Println("main ...")
	<-exit                          // 如果通道关闭，立即解除阻塞
	fmt.Println("main exit.")
}
