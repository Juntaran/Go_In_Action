/**
 * Author: Juntaran
 * Email:  Jacinthmail@gmail.com
 * Date:   2017/5/14 17:52
 */

package main

import (
	"sync"
	"math/rand"
	"fmt"
)

// 只有第一次调用该方法时传递给它的那个函数会执行

func main() {
	var count int
	var once sync.Once

	max := rand.Intn(100)
	for i := 0; i < max; i++ {
		once.Do(func() {
			count ++
		})
	}
	fmt.Printf("Count: %d.\n", count)
}
