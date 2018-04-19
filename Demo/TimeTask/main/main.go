/** 
  * Author: Juntaran 
  * Email:  Jacinthmail@gmail.com 
  * Date:   2018/4/17 11:05
  */

package main

import (
	"Go_In_Action/Demo/TimeTask"
	"fmt"
)

func test()  {
	fmt.Println("test")
}

func main() {
	//TimeTask.DoTimeTask(test, 2)
	TimeTask.DoTimeTask2(test, 2, 3)
}
