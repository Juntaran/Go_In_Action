/**
 * Author: Juntaran
 * Email:  Jacinthmail@gmail.com
 * Date:   2017/6/11 5:59
 */

package WorkPool

import (
	"fmt"
)

type Worker struct {
	ID 			int						// ID
	Work 		chan WorkRequest		// Work请求
	WorkerPool 	chan chan WorkRequest	// Worker池
	QuitChan 	chan bool				// 为true停止
}

// 创建一个worker，返回一个新的Worker对象
func NewWorker(id int, workerQueue chan chan WorkRequest) Worker {
	worker := Worker{
		ID: 		id,
		Work:		make(chan WorkRequest),
		WorkerPool: workerQueue,
		QuitChan: 	make(chan bool),
	}
	return worker
}

// 启动一个goroutine，for-select loop
func (w Worker) Start() {
	go func() {
		for {
			// 添加成员到worker队列
			w.WorkerPool <- w.Work
			select {
			case work := <- w.Work:
				// 接收到一个work请求
				fmt.Printf("worker: %d do it\n", w.ID)
				work.Execute(nil)
			case <- w.QuitChan:
				// 接收到停止要求
				fmt.Printf("worker: %d stopping\n", w.ID)
				return
			}
		}
	} ()
}

// 通知worker停止监听work请求，worker只有在完成当前工作后才会停止
func (w Worker) Stop() {
	go func() {
		w.QuitChan <- true
	} ()
}