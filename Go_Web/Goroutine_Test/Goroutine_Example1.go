package main

import (
	"runtime"
	"fmt"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()     // 把时间片让给别人，下次某个时候继续恢复执行该goroutine
		//runtime.GOMAXPROCS(4)   // 默认调度器只使用单线程，只实现了并发，GOMAXPROCS设置系统线程最大数量
		fmt.Println(s)
	}
}

func main() {
	go say("world")             // 开启一个新的Goroutine执行
	say("hello")                // 当前Goroutine执行
}