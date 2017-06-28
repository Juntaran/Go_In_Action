package main

import "fmt"

type Human struct {
	name    string
	age     int
	phone   string
}

type Employee struct {
	Human               // 匿名字段
	special string
	phone   string      // 重载
}

func main() {
	Bob := Employee{Human{"Bob", 18, "787878788"}, "Designer", "123456"}
	fmt.Println("Bob's work phone is ", Bob.phone)
	// 访问Human的phone
	fmt.Println("Bob's personal phone is ", Bob.Human.phone)
}