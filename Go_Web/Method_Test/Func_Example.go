package main

import "fmt"

type Rectangle struct {
	width, height   float64
}

func area(r Rectangle) float64 {
	return r.width * r.height
}

func main() {
	r1 := Rectangle{12, 2}
	r2 := Rectangle{9, 4}

	fmt.Println("r1 area is ", area(r1))
	fmt.Println("r2 area is ", area(r2))
}
