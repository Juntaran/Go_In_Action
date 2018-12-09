/**
  * Author: Juntaran
  * Email:  Jacinthmail@gmail.com
  * Date:   2017/4/7 19:28
  */

package main

import (
	"github.com/Juntaran/Go_In_Action/github.com/Juntaran/Go_In_Action/Chapter5/访问另一个包隐藏的标识符/Example2/entities"
	"fmt"
)

func main() {
	// 无法直接访问公开的结构类型中未公开的字段
	u := entities.User{
		Name:	"Bill",
		email:	"bill@email.com",		// Error！
	}
	fmt.Printf("User: %v\n", u)
}
