package main

import "fmt"

type Skills []string

type Human struct {
	name    string
	age     int
	weight  int
}

type Student struct {
	Human               // 匿名字段 struct
	Skills              // 匿名字段 自定义的string slice
	int                 // 匿名字段 内置类型int
	special string
}

func main() {
	// 初始化一个学生
	//Dendi := Student{Human{"Dendi", 18, 120}, special:"Computer"}
	Dendi := Student{Human:Human{"Dendi", 18, 120}, int:3, special:"CCC"}
	fmt.Println("His name is ", Dendi.name)
	fmt.Println("His special is ", Dendi.special)
	fmt.Println("His number is ", Dendi.int)

	// 修改skill字段
	Dendi.Skills = []string{"apple", "banana"}
	fmt.Println("His skills are ", Dendi.Skills)
	Dendi.Skills = append(Dendi.Skills, "dota2", "invoker")
	fmt.Println("His skills are ", Dendi.Skills)

	// 修改匿名内置类型字段
	Dendi.int = 7
	fmt.Println("His number is ", Dendi.int)
}