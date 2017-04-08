/**
  * Author: Juntaran
  * Email:  Jacinthmail@gmail.com
  * Date:   2017/4/7 19:56
  */

package main

import "fmt"

// 定义一个用户类型
type user struct {
	name 	string
	email 	string
}

// notify使用值接收者实现了一个方法
func (u user) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n", u.name, u.email)
}

// changeEmail使用指针接收者实现了一个方法
func (u *user) changeEmail(email string) {
	u.email = email
}

func main() {
	bill := user{
		"Bill",
		"bill@email.com",
	}
	bill.notify()

	lisa := &user{
		"Lisa",
		"lisa@email.com",
	}
	lisa.notify()

	/*
		Go语言会自动调整
		无论是值接收方法 还是 指针接收方法
		都可以直接使用
	*/

	bill.changeEmail("bill@new.com")
	bill.notify()

	lisa.changeEmail("lisa@new.com")
	lisa.notify()
}