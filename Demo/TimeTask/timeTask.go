/** 
  * Author: Juntaran 
  * Email:  Jacinthmail@gmail.com 
  * Date:   2018/4/17 11:00
  */

package TimeTask

import (
	"time"
	"fmt"
)

// 接收函数，定时处理
func DoTimeTask(f func(), sec int64)  {
	// 定时器 根据间隔时间 sec 执行函数 f()
	ticker := time.NewTicker(time.Second * time.Duration(sec))
	go func() {
		for _ = range ticker.C {
			f()
		}
	}()
	ch := make(chan int)
	<- ch
}

// 接收函数，间隔执行，after 秒后停止
func DoTimeTask2(f func(), interval, after time.Duration)  {
	timeout := time.After(after * time.Second)

	t := time.NewTicker(interval * time.Second)
	done := make(chan bool, 1)
	fmt.Println("执行第 1 次")
	f()
	go func() {
		i := 2
		for {
			select {
			case <-t.C:
				fmt.Println("执行第", i, "次")
				f()
				i++
			case <-timeout:
				close(done)
				return
			}
		}
	}()
	<-done
}