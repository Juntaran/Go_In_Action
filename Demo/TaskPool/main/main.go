/**
 * Author: Juntaran
 * Email:  Jacinthmail@gmail.com
 * Date:   2017/5/25 15:01
 */

package main

import (
	"fmt"
	"github.com/Juntaran/Go_In_Action/Demo/TaskPool"
	"time"
)

func test(i int, test string)  {
	fmt.Println("hahaha", i, test)
}

func main() {
	task_pool := TaskPool.WorkPool{}
	task_pool.InitPool(5)
	for i := 0; i < 100; i++ {
		task := TaskPool.Task{M_func: test}
		task.M_args = append(task.M_args, i)
		task.M_args = append(task.M_args, "test")
		task_pool.AddTask(task)
	}
	// 关闭任务池
	task_pool.ClosePool()

	time.Sleep(5 * time.Second)
	fmt.Println("Test Done!")
}
