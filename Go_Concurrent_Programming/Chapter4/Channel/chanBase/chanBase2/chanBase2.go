/**
  * Author: Juntaran
  * Email:  Jacinthmail@gmail.com
  * Date:   2017/5/10 23:42
  */

package main

import (
	"fmt"
	"time"
)

// 与chanBase1.go的区别在于用单向通道约束了用于发送或者接收的函数

// strChan1的作用是发送端告知接收端可以开始接收了
// strChan2的作用是保持主函数阻塞，避免两个goroutine没完成就已经退出了

var strChan = make(chan string, 3)

func receive(strChan <-chan string, syncChan1 <-chan struct{}, syncChan2 chan<- struct{})  {
	<-syncChan1
	fmt.Println("Received a sync signal and wait a second... [receiver]")
	time.Sleep(time.Second)

	for {
		if elem, ok := <-strChan; ok {
			fmt.Println("Received:", elem, "[receiver]")
		} else {
			break
		}
	}
	fmt.Println("Stopped. [receiver]")
	syncChan2 <- struct{}{}
}

func send(strChan chan<- string, syncChan1 chan<- struct{}, syncChan2 chan<- struct{})  {
	for _, elem := range []string{"a", "b", "c", "d"} {
		strChan <- elem
		fmt.Println("Sent:", elem, "[sender]")
		if elem == "c" {
			syncChan1 <- struct{}{}
			fmt.Println("Sent a sync signal. [sender]")
		}
	}
	fmt.Println("Wait 2 seconds... [sender]")
	time.Sleep(time.Second * 2)
	close(strChan)
	syncChan2 <- struct{}{}
}


func main() {
	syncChan1 := make(chan struct{}, 1)
	syncChan2 := make(chan struct{}, 2)

	// 接收操作
	go receive(strChan, syncChan1, syncChan2)
	// 发送操作
	go send(strChan, syncChan1, syncChan2)

	<- syncChan2
	<- syncChan2
}