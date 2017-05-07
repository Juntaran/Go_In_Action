package main

import (
	"sync"
	"math/rand"
	"time"
	"fmt"
)

/*
	无缓冲通道是指在接收前没有任何能力保存任何值的通道
	需要发送goroutine和接收goroutine同时准备好，才能完成发送、接收操作
	否则，通道会导致先执行操作的goroutine阻塞
*/

/*
	一个发送语句将一个值从一个goroutine通过channel发送到另一个执行接收操作的goroutine。
	发送和接收两个操作都是用<-运算符。在发送语句中，<-运算符分割channel和要发送的值。
	在接收语句中，<-运算符写在channel对象之前。一个不使用接收结果的接收操作也是合法的。
*/

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())		// 根据当前时间纳秒设置随机数种子
}

// 模拟一个选手打网球
func player(name string, court chan int)  {
	defer wg.Done()

	for {
		// 等待球被击打过来
		ball, ok := <-court

		if !ok {
			// 如果通道关闭，胜利
			fmt.Printf("Player %s Won\n", name)
			return
		}

		// 选随机数，根据随机数判断是否丢球
		n := rand.Intn(100)								// Intn(100)在(0,100)随机取一个非负整数
		if n % 13 == 0 {
			fmt.Printf("Player %s Missed\n", name)
			// 关闭通道，失败
			close(court)
			return
		}

		// 显示击球数，击球数+1
		fmt.Printf("Player %s Hit %d\n", name, ball)
		ball ++

		// 将球打向对手
		court <- ball
	}
}

func main() {
	// 创建一个无缓冲通道
	court := make(chan int)

	wg.Add(2)

	go player("A", court)
	go player("B", court)

	// 发球
	court <- 1

	// 等待结束
	wg.Wait()
}
