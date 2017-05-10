/**
  * Author: Juntaran
  * Email:  Jacinthmail@gmail.com
  * Date:   2017/5/10 16:27
  */

package main

import (
	"fmt"
	"time"
)

var mapChan = make(chan map[string]int, 1)

// mapChan的元素类型map属于引用类型，接收方对元素值副本的修改会影响发送方持有的源值

func main() {
	syncChan := make(chan struct{}, 2)

	go func() {
		// 接收操作
		for {
			if elem, ok := <-mapChan; ok {
				elem["count"] ++
			} else {
				break
			}
		}
		fmt.Println("Stopped. [receiver]")
		syncChan <- struct{}{}
	} ()

	go func() {
		// 发送操作
		countMap := make(map[string]int)
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
