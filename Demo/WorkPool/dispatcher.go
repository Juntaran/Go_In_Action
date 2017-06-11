/**
 * Author: Juntaran
 * Email:  Jacinthmail@gmail.com
 * Date:   2017/6/11 6:25
 */

package WorkPool

import "fmt"

type WorkerPoolType chan chan WorkRequest

var WorkerPool WorkerPoolType

func StartDispatcher(nworkers int)  {
	// 初始化通道
	WorkerPool = make(WorkerPoolType, nworkers)

	// 创建workers
	for i := 0; i < nworkers; i++ {
		fmt.Println("Starting worker", i + 1)
		worker := NewWorker(i+1, WorkerPool)
		worker.Start()
	}

	go func() {
		for {
			select {
			case work := <-WorkQueue:
				fmt.Println("Received work request.")
				go func() {
					worker := <-WorkerPool
					fmt.Println("Dispatching work request")
					worker <- work
				} ()
			}
		}
	} ()
}