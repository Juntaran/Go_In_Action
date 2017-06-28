package main

import "fmt"

func Fibonacci(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int, 10)
	go Fibonacci(cap(c), c)
	//fmt.Println(cap(c))
	for i := range c {          // 不断读取channel里的数据，直到channel被关闭
		fmt.Println(i)
	}
}