package main

import "fmt"

type person struct {
	name    string
	age     int
}

func Older(p1, p2 person) (person, int) {
	if p1.age > p2.age {
		return p1, p1.age-p2.age
	}
	return p2, p2.age-p1.age
}

func main() {
	var Tom person
	Tom.name, Tom.age = "Tom", 18
	Bob := person{name:"Bob", age:25}
	Paul:= person{"Paul", 43}

	TB_Older, TB_diff := Older(Tom, Bob)
	TP_Older, TP_diff := Older(Tom, Paul)
	BP_Older, BP_diff := Older(Bob, Paul)

	fmt.Println(TB_Older.name, TB_diff)
	fmt.Println(TP_Older.name, TP_diff)
	fmt.Println(BP_Older.name, BP_diff)

}
