/**
  * Author: Juntaran
  * Email:  Jacinthmail@gmail.com
  * Date:   2017/4/7 19:31
  */

package entities

// user是一个未公开的结构类型
// 未公开的内部类型，代码无法直接通过结构字面量的方式初始化该内部类型
type user struct {
	Name 	string
	Email	string
	test 	string
}

// Admin为一个公开的类型
type Admin struct {
	user			// 嵌入的内部类型未公开
	Rights	int
}