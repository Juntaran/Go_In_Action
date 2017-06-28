package main

import (
	"reflect"
	"fmt"
)

type X int
type Y int

func main() {
	var a X = 100
	t := reflect.TypeOf(a)
	v := reflect.ValueOf(a)
	fmt.Println(t.Name(), t.Kind())
	fmt.Println(v)
}
