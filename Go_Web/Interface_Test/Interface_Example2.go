package main

import "fmt"

func main() {
	// 定义a为空接口
	// 空interface可以存储任意类型的数值
	var a interface{}
	i := 5
	s := "Hello, World"

	// a可以存储任意类型的数值
	fmt.Println(a)
	a = i
	fmt.Println(a)
	a = s
	fmt.Println(a)
}
