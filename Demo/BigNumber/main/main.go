/** 
  * Author: Juntaran 
  * Email:  Jacinthmail@gmail.com 
  * Date:   2017/9/6 17:30 
  */

package main

import (
	"Go_In_Action/Demo/BigNumber"
	"fmt"
)

func main() {
	a := "-1112357657"
	b := "+45657768"

	c, err1 := BigNumber.BigAdd(a, b)
	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Println(c)
	}

	ret, err2 := BigNumber.BigReduce(c, a)
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println(ret)
	}
}
