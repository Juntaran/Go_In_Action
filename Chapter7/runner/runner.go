/**
  * Author: Juntaran
  * Email:  Jacinthmail@gmail.com
  * Date:   2017/4/11 02:13
  */

package runner

import (
	"os"
	"time"
	"errors"
	"os/signal"
)

// 利用通道监视程序的执行时间，如果时间太长则可以终止
// 这个程序可能会作为cron作业执行

// Runner在给定的超时时间内执行一组任务
// 在操作系统发送中断信号时结束任务
type Runner struct {
	// interrupt 通道报告从操作系统发送的信号
	interrupt chan os.Signal
	// complete  通道报告处理任务已经完成
	complete chan error
	// timeout   报告处理任务已经超时
	timeout <-chan time.Time
	// tasks 持有一组以索引顺序依次执行的函数
	tasks []func(int)
}

// ErrTimeout在任务超时返回
// New returns an error that formats as the given text.
var ErrTimeout = errors.New("received timeout")

// ErrInterrupt在接收到操作系统事件时返回
var ErrInterrupt = errors.New("received interrupt")

// New返回一个新的准备使用的Runner
func New(duration time.Duration) *Runner {
	return &Runner{
		interrupt:	make(chan os.Signal, 1),
		complete: 	make(chan error),
		timeout: 	time.After(duration),
	}
}

// Add把一个任务附加到Runner上
// 这个任务是一个接收到一个int类型的ID作为参数的函数
func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

// Start执行所有任务，并监视通道事件
func (r *Runner) Start() error {
	// 接收所有的中断信号
	signal.Notify(r.interrupt, os.Interrupt)
	// 不同的goroutine执行不同的任务
	go func() {
		r.complete <- r.run()
	}()
	
	select {
	// 当任务处理完成时发出的信号
	case err := <-r.complete:
		return err
	// 当任务处理超时发出的信号
	case <-r.timeout:
		return ErrTimeout
	}
}

// run执行每一个已经注册的服务
func (r *Runner) run() error {
	for id, task := range r.tasks {
		// 检测到操作系统中断信号
		if r.gotInterrupt() {
			return ErrInterrupt
		}
		// 执行已经注册的任务
		task(id)
	}
	return nil
}

// gotInterrupt验证是否接收到了中断信号
func (r *Runner) gotInterrupt() bool {
	select {
	// 当中断事件被触发时发出的信号
	case <-r.interrupt:
		// 停止接收后续的任何信号
		signal.Stop(r.interrupt)
		return true

	// 继续正常运行
	default:
		return false
	}
}