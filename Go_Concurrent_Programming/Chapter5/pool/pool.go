/**
 * Author: Juntaran
 * Email:  Jacinthmail@gmail.com
 * Date:   2017/5/14 18:02
 */

package main

import (
	"runtime/debug"
	"sync/atomic"
	"sync"
	"fmt"
	"runtime"
)

/*
	sync.Pool 类型可以看作存放临时值得容器，并发安全
	优点一：临时对象池可以把由其中的对象值产生的存储压力进行分摊
	优点二：对垃圾回收友好，永远不会显式从池中取走某个对象值，该对象值也不会永远在临时池中
		   它的生命周期取决于垃圾回收任务的下一次执行时间
*/

func main() {
	// 禁用GC，并保证在main函数执行结束前恢复GC
	defer debug.SetGCPercent(debug.SetGCPercent(-1))
	var count int32

	newFunc := func() interface{} {
		return atomic.AddInt32(&count, 1)
	}
	pool := sync.Pool{New: newFunc}

	// New字段值作用
	v1 := pool.Get()
	fmt.Printf("Value 1: %v\n", v1)

	// 临时对象池的存取
	pool.Put(10)
	pool.Put(11)
	pool.Put(12)
	v2 := pool.Get()
	fmt.Printf("Value 2: %v\n", v2)

	// 垃圾回收对临时对象池的影响
	debug.SetGCPercent(100)
	runtime.GC()
	v3 := pool.Get()
	fmt.Printf("Value 3: %v\n", v3)
	pool.New = nil
	v4 := pool.Get()
	fmt.Printf("Value 4: %v\n", v4)
}
