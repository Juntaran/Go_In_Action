package main

import (
	"time"
	"fmt"
)

// 一个简易动画
func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

// 斐波那契数列
func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}

func main() {
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib(n)
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

/*
	主函数返回时，所有的goroutine都会直接打断，程序退出
	除了从主函数退出或者直接退出程序之外，
	没有其他的编程方法能够让一个goroutine来打断另一个的执行
	但是可以通过goroutine之间的通信来让一个goroutine请求去请求其它的goroutine
	之后让自己结束执行
*/