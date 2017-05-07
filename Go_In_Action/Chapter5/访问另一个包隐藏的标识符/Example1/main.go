/**
  * Author: Juntaran
  * Email:  Jacinthmail@gmail.com
  * Date:   2017/4/7 19:15
  */

package main

import (
	"fmt"
	"Go_In_Action/Go_In_Action/Chapter5/访问另一个包隐藏的标识符/Example1/counters"
)

// 访问另一个包的未公开标识符的值
func main() {
	// 使用counters包公开的New函数来创建一个未公开的类型的变量
	// 使用New作为工厂函数的名字是Go开发习惯
	counter := counters.New(10)
	fmt.Printf("Counter: %d\n", counter)
}
