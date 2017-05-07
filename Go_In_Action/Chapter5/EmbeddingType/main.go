/**
  * Author: Juntaran
  * Email:  Jacinthmail@gmail.com
  * Date:   2017/4/7 18:53
  */

package main

import "fmt"

/*
	嵌入类型能够帮助用户扩展或修改已有类型的行为
	嵌入类型是把已有的类型直接声明在新的结构类型里
	被嵌入的类型被称为新的外部类型的内部类型
*/

// 定义一个用户类user
type user struct {
	name 	string
	email	string
}

// 定义一个管理员类admin
type admin struct {
	user			// 这里的user是一个嵌入类型
//	user	user  	// 这样的user不是嵌入类型
	level	string
}

// notify实现了一个可以通过user类型值的指针
func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
}

func main() {
	// 创建一个管理员用户
	a := admin{
		user:	user{
			name:	"john smith",
			email:	"john@yahoo.com",
		},
		level:	"super",
	}
	// 可以直接访问内部类型的方法
	a.user.notify()

	// 内部类型的方法也能被提升到外部类型
	a.notify()
}