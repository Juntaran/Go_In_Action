/**
 * Author: Juntaran
 * Email:  Jacinthmail@gmail.com
 * Date:   2017/5/25 14:51
 */

package TaskPool

import (
	"reflect"
	"time"
)

type Task struct {
	M_func 	interface{}
	M_args	[]interface{}
}

func (task *Task) Run() {
	go func() {
		f := reflect.ValueOf(task.M_func)
		if len(task.M_args) != f.Type().NumIn() {
			return
		}
		in := make([]reflect.Value, len(task.M_args))
		for k, param := range task.M_args {
			in[k] = reflect.ValueOf(param)
		}
		f.Call(in)
	} ()
}

type WorkPool struct {
	TaskChannel		chan Task
	QuitChan		chan int		// 终止通道
}

// size为缓存大小
func (pool *WorkPool) InitPool(size int) {
	pool.TaskChannel = make(chan Task, size)
	pool.QuitChan = make(chan int)
	go func() {
	DONE:
		for {
			select {
			case task := <-pool.TaskChannel:
				task.Run()
			case <-pool.QuitChan:
				break DONE
			}
		}
	} ()
}

func (pool *WorkPool) ClosePool() {
	pool.QuitChan <- 1
}

// 同步阻塞方式添加任务
func (pool *WorkPool) AddTask(task Task) {
	pool.TaskChannel <- task
}

// 非阻塞方式添加任务，time为超时时间，单位毫秒
func (pool *WorkPool) AddTaskSync(task Task, millitime int) bool {
	res := false
	go func(res bool) {
		select {
		case pool.TaskChannel <- task:
			res = true
		case <-time.After(time.Millisecond * time.Duration(millitime)):
			res = false
		}
	} (res)
	return res
}