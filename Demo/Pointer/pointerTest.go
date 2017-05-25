/**
 * Author: Juntaran
 * Email:  Jacinthmail@gmail.com
 * Date:   2017/5/25 15:14
 */

package main

import "fmt"

type T struct {
	name	string
	value 	int
}

type Vertex struct {
	X 		int
	Y 		string
}

func main() {
	var a int = 20
	var s string = "a"
	fmt.Printf("a变量存储地址：%x\n", &a)
	fmt.Printf("s变量存储地址：%x\n\n", &s)

	var p *int		// int类型的指针
	fmt.Printf("p变量存储地址：%x\n\n", &p)
	// 使用指针访问值
	i := 10
	p = &i
	fmt.Printf("i变量存储地址：%x\n", &i)
	fmt.Printf("p变量存储地址：%x\n", &p)
	fmt.Printf("*p变量的值为：%d\n\n", *p)

	*p = 20
	fmt.Printf("i变量的值为：%d\n", i)
	fmt.Printf("i变量存储地址：%x\n", &i)
	fmt.Printf("p变量存储地址：%x\n", &p)
	fmt.Printf("*p变量的值为：%d\n\n", *p)

	var b int
	var ptr *int
	var pptr **int
	b = 10
	ptr = &b
	pptr = &ptr

	fmt.Printf("b变量的值为：%d\n", b)
	fmt.Printf("b变量存储地址为：%x\n", &b)
	fmt.Printf("指针变量*ptr的值为：%d\n", *ptr)
	fmt.Printf("指针变量*ptr的存储地址为：%x\n", &*ptr)
	fmt.Printf("二重指针变量*pptr的值为：%d\n", **pptr)
	fmt.Printf("二重指针变量*pptr的存储地址为：%x\n\n", &**pptr)

	var v = Vertex{
		X: 	1,
		Y: 	"s",
	}
	fmt.Println(v)
	fmt.Println(v.X)
	fmt.Println(v.Y)
	v.X = 10
	fmt.Println(v.X)

	pV := &v					// 一个指向结构体的指针pV
	pV.X = 33333333
	fmt.Println(v)
	fmt.Println((*pV).X)		// 可以用 (*pV).X 访问字段X
	fmt.Println(pV.X)			// 也可以隐式间接引用
}