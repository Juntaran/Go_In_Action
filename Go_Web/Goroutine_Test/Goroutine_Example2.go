package main

import (
	"time"
	"fmt"
)

var c int

func counter() int {
	c ++
	return c
}

func main() {
	a := 100
	go func(x, y int) {
		time.Sleep(time.Second * 3)         // goroutine在main逻辑之后执行
		fmt.Println("Go: ", x, y)
	} (a, counter())                        // 立即计算并复制参数
	a += 100
	fmt.Println("main: ", a, counter())
	time.Sleep(time.Second * 5)             //　等待goroutine结束
}