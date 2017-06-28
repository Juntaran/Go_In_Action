/**
 * Author: Juntaran
 * Email:  Jacinthmail@gmail.com
 * Date:   2017/6/28 10:12
 */

package main

import (
	"Go_In_Action/Demo/HeartBeat"
	"fmt"
	"time"
)

func main() {
	tm := HeartBeat.NewTaskMap("test")
	name, duration := "First", 1
	task, err := HeartBeat.NewTask(tm, name, duration)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(task.CreateTime)
		// 启动一个任务
		task.Start(tm, func() error {
			fmt.Println(name)
			return nil
		})
		time.Sleep(time.Second * 3)
		// 获取状态
		fmt.Println("获取任务状态：")
		fmt.Println(HeartBeat.GetActivity(tm))

		// 暂停任务
		fmt.Println("暂停任务：")
		if err := HeartBeat.PauseTask(tm, name); err != nil {
			fmt.Println(err)
		}
		time.Sleep(1)
		fmt.Println(HeartBeat.GetActivity(tm))

		// 启动任务
		fmt.Println("启动任务：")
		if err := HeartBeat.RunTask(tm, name); err != nil {
			fmt.Println(err)
		}
		fmt.Println(HeartBeat.GetActivity(tm))

		// 删除任务
		fmt.Println("删除任务：")
		if err := HeartBeat.DelTask(tm, name); err != nil {
			fmt.Println(err)
		}
		fmt.Println(HeartBeat.GetActivity(tm))
	}
}
