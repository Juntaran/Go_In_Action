/**
  * Author: Juntaran
  * Email:  Jacinthmail@gmail.com
  * Date:   2017/5/10 16:59
  */

package main

import "fmt"

// 调用close函数可以关闭一个通道，试图向一个已经关闭的通道发送元素值会让发送操作panic
// 在保证安全的前提下关闭通道
// 无论如何不要在接收方关闭通道，接收端通常无法判断发送端是否还有要发送的元素

func main() {

	dataChan  := make(chan int, 5)
	syncChan1 := make(chan struct{}, 1)
	syncChan2 := make(chan struct{}, 2)

	go func() {
		// 接收操作
		<- syncChan1
		for {
			// ok可以判断通道是否关闭
			if elem, ok := <-dataChan; ok {
				fmt.Printf("Received: %d [receiver]\n", elem)
			} else {
				break
			}
		}
		fmt.Println("Done. [receiver]")
		syncChan2 <- struct{}{}
	} ()

	go func() {
		// 发送操作
		for i := 0; i < 5; i++ {
			dataChan <- i
			fmt.Printf("Sent: %d [sender]\n", i)
		}
		close(dataChan)
		syncChan1 <- struct{}{}
		fmt.Println("Done. [sender]")
		syncChan2 <- struct{}{}
	} ()

	<- syncChan2
	<- syncChan2
}
