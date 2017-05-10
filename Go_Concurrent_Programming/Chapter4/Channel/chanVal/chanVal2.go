/**
  * Author: Juntaran
  * Email:  Jacinthmail@gmail.com
  * Date:   2017/5/10 16:38
  */

package main

import (
	"fmt"
	"time"
)

// Counter计数器
type Counter struct {
	count int
}

// 利用结构体则接收方对元素值的修改不会影响发送方的源值
// 如果在mapChan传递的是 map[string]*Counter，则仍会修改发送方的源值

var mapChan = make(chan map[string]Counter, 1)
//var mapChan = make(chan map[string]*Counter, 1)

func (counter *Counter) String() string {
	return fmt.Sprintf("{count: %d}", counter.count)
}

func main() {
	syncChan := make(chan struct{}, 2)

	go func() {
		// 接收操作
		for {
			if elem, ok := <- mapChan; ok {
				counter := elem["count"]
				counter.count ++
			} else {
				break
			}
		}
		fmt.Println("Stopped. [receiver]")
		syncChan <- struct{}{}
	} ()

	go func() {
		// 发送操作
		countMap := map[string]Counter{
			"count" : Counter{},
		}
		//countMap := map[string]*Counter{
		//	"count" : &Counter{},
		//}

		for i := 0; i < 5; i++ {
			mapChan <- countMap
			time.Sleep(time.Millisecond)
			fmt.Printf("The count map: %v. [sender]\n", countMap)
		}
		close(mapChan)
		syncChan <- struct{}{}
	} ()

	<- syncChan
	<- syncChan
}