/**
 * Author: Juntaran
 * Email:  Jacinthmail@gmail.com
 * Date:   2017/4/11 02:35
 */

package main

import (
	"Go_In_Action/Go_In_Action/Chapter7/runner"
	"log"
	"os"
	"time"
)

// timeout规定必须在多少秒内完成任务
const timeout = 3 * time.Second

// createTask 返回一个根据id休眠指定秒数的示例任务
func createTask() func(int) {
	return func(id int) {
		log.Printf("Processor - Task #%d.", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}

func main() {
	log.Println("Starting work.")

	// 为本次执行分配超时时间
	r := runner.New(timeout)

	// 加入要执行的任务
	r.Add(createTask(), createTask(), createTask())

	// 执行任务并处理结果
	if err := r.Start(); err != nil {
		switch err {
		case runner.ErrTimeout:
			log.Println("Terminating due to Timeout.")
			os.Exit(1)
		case runner.ErrInterrupt:
			log.Println("Terminating due to Interrupt.")
			os.Exit(2)
		}
	}
	log.Println("Process ended.")
}
