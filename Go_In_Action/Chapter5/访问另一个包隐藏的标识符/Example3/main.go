/**
  * Author: Juntaran
  * Email:  Jacinthmail@gmail.com
  * Date:   2017/4/7 19:33
  */

package main

import (
	"Go_In_Action/Go_In_Action/Chapter5/访问另一个包隐藏的标识符/Example3/entities"
	"fmt"
)

func main() {
	// 创建entities包里的Admin类型的值
	a := entities.Admin{
		Rights:	10,
	}
	// 设置未公开的内部类型的公开字段的值
	a.Name  = "Bill"
	a.Email = "bill@example.com"

	var b entities.Admin
	b.Name   = "Juntaran"
	b.Email  = "jacinthmail@gmail.com"
	b.Rights = 1

	fmt.Printf("User1: %v\n", a)
	fmt.Printf("User2: %v\n", b)
}
