/**
  * Author: Juntaran
  * Email:  Jacinthmail@gmail.com
  * Date:   2017/5/11 15:00
  */

package LoadGen

import (
	"time"
	"github.com/Juntaran/Go_In_Action/Go_Concurrent_Programming/Chapter4/LoadGen/lib"
	"context"
)

// 载荷发生器的实现类型
type myGenerator struct {
	timeoutNS		time.Duration			// 响应超时时间ns
	lps 			uint32					// 每秒载荷量
	durationNS 		time.Duration			// 负载持续时间ns
	tickets 		lib.GoTickets			// Goroutine票池
	ctx 			context.Context			// 上下文
	cancleFunc 		context.CancelFunc		// 取消函数


	resultCh        chan *lib.CallResult	// 调用结果通道
}
