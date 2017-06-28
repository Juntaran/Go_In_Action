package main

import "fmt"

// 必须使用make创建channel

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum                    // 把sum发送到channel c
}

func main() {
	a := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)         // 创建一个channel，非缓存类型
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <- c, <-c
	fmt.Println(x, y, x+y)
}
