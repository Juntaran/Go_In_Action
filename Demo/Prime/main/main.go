/**
  * Author: Juntaran
  * Email:  Jacinthmail@gmail.com
  * Date:   2017/5/9 20:07
  */

package main

import "Go_In_Action/Demo/Prime"

func main() {
	origin, wait := make(chan int), make(chan struct{})
	Prime.Processor(origin, wait)
	for num := 2; num < 10000; num++ {
		origin <- num
	}
	close(origin)
	<- wait
}
