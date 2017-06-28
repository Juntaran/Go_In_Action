package main

import "fmt"

type Human struct {
	name    string
	age     int
	phone   string
}

type Student struct {
	Human
	school  string
	loan    float64
}

type Employee struct {
	Human
	company string
	salary  float64
}

// Human对象实现SayHi方法
func (h Human) SayHi() {
	fmt.Printf("Hi, I am %s, you can call me on %s\n", h.name, h.phone)
}

// Human对象实现Sing方法
func (h Human) Sing(lyric string) {
	fmt.Println("La la, la la la,", lyric)
}

// Human对象实现Guzzle方法
func (h Human) Guzzle(beerStein string) {
	fmt.Println("Guzzle Guzzle Guzzle...", beerStein)
}

// Employee重载Human的SayHi方法
func (e Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s\n", e.name, e.company)
}

// Student实现BorrowMoney方法
func (s Student) BorrowMoney(money float64) {
	s.loan += money
}

// Employee实现SpendSalary方法
func (e Employee) SpendSalary(money float64) {
	e.salary -= money
}

// 定义Interface
// Men被Human Student 和 Employee实现，因为这三个类型都实现了这两个方法
type Men interface {
	SayHi()
	Sing(lyric string)
}

func main() {
	Mike := Student {Human{"Mike", 25, "2222222"}, "MIT",    0.00}
	Paul := Student {Human{"Paul", 26, "1111111"}, "HEU",    100}
	Sam  := Employee{Human{"Sam",  36, "3333333"}, "Golang", 1000}
	Tom  := Employee{Human{"Tom",  45, "4444444"}, "Java",   5000}

	// 定义Men类型的变量i
	var i Men

	// i能存储Student
	i = Mike
	fmt.Println("This is Mike, a Student:")
	i.SayHi()
	i.Sing("Good Good Study")

	// i也能存储Employee
	i = Tom
	fmt.Println("This is Tom, an Employee:")
	i.SayHi()
	i.Sing("Day Day Up")

	// 定义了Slice Men
	fmt.Println("A slice of Men")
	x := make([]Men, 3)
	// 他们都是不同类型的元素，但他们实现了interface同一接口
	x[0], x[1], x[2] = Paul, Sam, Mike

	for _, value := range x{
		value.SayHi()
	}
}