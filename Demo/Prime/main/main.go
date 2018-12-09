/**
  * Author: Juntaran
  * Email:  Jacinthmail@gmail.com
  * Date:   2017/5/9 20:07
  */

package main

import (
	"github.com/Juntaran/Go_In_Action/Demo/Prime"
	"fmt"
)

func main() {
	ret, count := Prime.GetPrime(10000)
	fmt.Printf("Prime count is %d\n", count)
	fmt.Println(ret)
}
