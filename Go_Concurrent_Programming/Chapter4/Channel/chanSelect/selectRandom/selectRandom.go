/**
  * Author: Juntaran
  * Email:  Jacinthmail@gmail.com
  * Date:   2017/5/11 00:33
  */

package main

import "fmt"

/*
	在执行select语句时，系统自上而下判断每个case的发送或接收操作是否可以立即进行
	这里的立即进行是指当前goroutine不会因此操作而被阻塞
	当一个case被选中，会执行该case语句
	如果同时有多个case满足条件，系统会通过一个伪随机算法选中一个case
*/

func main() {
	chanCap := 5
	intChan := make(chan int, chanCap)
	for i := 0; i < chanCap; i++ {
		select {
		case intChan <- 1:
		case intChan <- 2:
		case intChan <- 3:
		}
	}
	for i := 0; i < chanCap; i++ {
		fmt.Printf("%d\n", <-intChan)
	}
}
