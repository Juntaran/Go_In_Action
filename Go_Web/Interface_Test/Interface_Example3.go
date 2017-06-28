package main

import (
	"strconv"
	"fmt"
)

/*
// fmt的源码，也就是说任何实现了String方法的类型都能作为参数被fmt.Println()调用
type Stinger interface {
	String()    string
}
*/

type Human struct {
	name    string
	age     int
	phone   string
}

// 通过这个方法Human实现了fmt.Stringer
func (h Human) String() string {
	return "_" + h.name + "-" + strconv.Itoa(h.age) + "-years-" + h.phone + "_"
}

func main() {
	Bob := Human{"Bob", 39, "0451-88888888"}
	fmt.Println("This Human is : ", Bob)
}