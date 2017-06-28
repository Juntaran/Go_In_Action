package main

import "fmt"

type Human struct {
	name    string
	age     int
	weight  int
}

type Student struct {
	Human           // 匿名字段，默认Student包含Human所有字段
	special string
}

func main() {
	// 初始化一个学生
	Dendi := Student{Human{"Dendi", 25, 120}, "Computer Science"}

	fmt.Println("His name is ", Dendi.name)
	fmt.Println("His Special is ", Dendi.special)

	Dendi.special = "AI"
	Dendi.age = 18
	fmt.Println("His Special is ", Dendi.special)
	fmt.Println("His age is ", Dendi.age)
}