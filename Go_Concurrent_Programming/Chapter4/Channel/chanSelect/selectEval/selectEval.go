/**
  * Author: Juntaran
  * Email:  Jacinthmail@gmail.com
  * Date:   2017/5/11 00:15
  */

package main

import "fmt"

var intChan1 chan int
var intChan2 chan int
var channels = []chan int{intChan1, intChan2}

var numbers = []int{1, 2, 3, 4, 5, 6}

func getNumber(i int) int {
	fmt.Printf("numbers[%d]\n", i)
	return numbers[i]
}

func getChan(i int) chan int {
	fmt.Printf("channels[%d]\n", i)
	return channels[i]
}

func main() {
	select {
	case getChan(0) <- getNumber(0):
		fmt.Println("1st case is selected.")
	case getChan(1) <- getNumber(1):
		fmt.Println("2nd case is selected.")
	default:
		fmt.Println("Default!")
	}
}
