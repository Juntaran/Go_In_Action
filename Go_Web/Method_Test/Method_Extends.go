package main

import "fmt"

type Human struct {
	name    string
	age     int
	phone   string
}

type Student struct {
	Human   // 匿名字段
	school  string
}

type Employee struct {
	Human   // 匿名字段
	company string
}

// 在human定义一个method
func (h Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

func main() {
	Mark := Student{Human{"Mark", 25, "136845"}, "MIT"}
	Sam  := Employee{Human{"Sam", 45, "11111"}, "HEU"}

	Mark.SayHi()
	Sam.SayHi()
}