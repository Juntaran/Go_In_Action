package main

import "fmt"


/*
	ch := make(chan type, value)
	value == 0      // 无缓冲
	value > 0       // 缓冲（非阻塞，直到value个元素）

*/

func main() {
	c := make(chan int, 2)
	c <- 1
	c <- 2
	fmt.Println(<-c)
	fmt.Println(<-c)
}
