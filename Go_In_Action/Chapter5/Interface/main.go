/**
  * Author: Juntaran
  * Email:  Jacinthmail@gmail.com
  * Date:   2017/4/7 04:53
  */

package main

import (
	"fmt"
)

// 一个用户类型user
type user struct {
	name 	string
	email	string
}

// 一个管理员类型admin
type admin struct {
	user			// 这里user是一个内部类
	power	bool
}

// 通知接口notifier
type notifier interface {
	notify()
	test()
}

// nitify方法的user实现
// notify方法使用值接收者实现了notifier接口
func (u user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
}

// notify方法的admin实现
// notify方法使用指针接收者实现了notifier接口
func (a *admin) notify() {
	fmt.Printf("Sending admin email to %s<%s>, admin power is %v\n", a.user.name, a.user.email, a.power)
}

func (a *admin) test() {
	fmt.Println("admin test")
}

func (u user) test() {
	fmt.Println("user test")
}

// sendNotification接受了一个实现了notifier接口的值
func sendNotification(n notifier)  {
	n.notify()
	n.test()
}

func main() {
	// 创建一个user，传给sendNotification
	bill := user{
		"Bill",
		"bill@email.com",
	}
	sendNotification(bill)

	// 创建一个admin，传给sendNotification
	lisa := admin{
		user{
			name: "Lisa",
			email:"lisa@email.com",
		},
		true,
	}
	sendNotification(&lisa)

	// 因为是使用指针接收者实现的接口func (a *admin) notify() {}
	// 所以应该使用sendNotification(&lisa) 而不是 sendNotification(lisa)

	bill.notify()

	lisa.notify()
	lisa.user.notify()
}