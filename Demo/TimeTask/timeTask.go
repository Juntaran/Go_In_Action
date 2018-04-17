/** 
  * Author: Juntaran 
  * Email:  Jacinthmail@gmail.com 
  * Date:   2018/4/17 11:00
  */

package TimeTask

import (
	"time"
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
