/**
 * Author: Juntaran
 * Email:  Jacinthmail@gmail.com
 * Date:   2017/6/28 9:43
 */

package HeartBeat

import (
	"sync"
	"time"
	"errors"
	"runtime"
)

const (
	Running = "running"
	Pause = "pause"
	Stop = "stop"
)

// 任务结构体
type Task struct {
	Name 		string
	Status 		string
	Duration 	int
	Chan 		chan string				// 接收信号
	CreateTime	time.Time
	lock 		*sync.Mutex
}

// 任务索引结构
type TaskMap struct {
	Name		string					// 可能有多种任务，每种对应一个name
	task 		map[string]*Task
	lock 		*sync.Mutex
}

// 创建TaskMap
func NewTaskMap(name string) *TaskMap {
	taskMap := new(TaskMap)
	taskMap.Name = name
	taskMap.task = make(map[string]*Task)
	taskMap.lock = new(sync.Mutex)
	return taskMap
}

// 添加索引
func addMaps(taskMap *TaskMap, name string, task *Task)  {
	taskMap.lock.Lock()
	defer taskMap.lock.Unlock()
	{
		taskMap.task[name] = task
	}
}

// 删除索引
func delMap(taskMap *TaskMap, name string)  {
	taskMap.lock.Lock()
	defer taskMap.lock.Unlock()
	{
		delete(taskMap.task, name)
	}
}

// 查找索引
func getMap(taskMap *TaskMap, name string) *Task {
	if taskMap == nil {
		return nil
	}
	return taskMap.task[name]
}

// 创建一个任务，该任务默认运行
func NewTask(taskMap *TaskMap, name string, duration int) (*Task, error) {
	if name == "" {
		return nil, errors.New("Error: Name cannot be Empty")
	}
	if getMap(taskMap, name) != nil {
		return nil, errors.New("Error: Name Conflict")
	}
	task := &Task{
		Name: 			name,
		Status: 		Running,
		Duration: 		duration,
		Chan: 			make(chan string),
		CreateTime:		time.Now(),
	}
	addMaps(taskMap, name, task)
	return task, nil
}

func run(taskMap *TaskMap, task *Task, f func() error)  {
	timer := time.NewTicker(time.Duration(task.Duration) * time.Second)
	for {
		select {
		case <- timer.C:
			if task.Status == Pause {
				runtime.Gosched()
				continue
			}
			if err := f(); err != nil {
				timer.Stop()
				return
			}
		case status, ok := <-getMap(taskMap, task.Name).Chan:
			if !ok {
				if ret := getMap(taskMap, task.Name); ret != nil {
					close(ret.Chan)
				}
			}
			switch status {
			case Stop:
				timer.Stop()
				return
			case Running:
				task.Status = Running
			case Pause:
				task.Status = Pause
			}
		}
	}
}

func (task *Task) Start(taskMap *TaskMap, f func() error) {
	go run(taskMap, task, f)
}

// 获取状态
func GetActivity(taskMap *TaskMap) (ret []interface{}) {
	for _, k := range taskMap.task {
		if k != nil {
			dict := make(map[string]interface{})
			dict["Name"] = k.Name
			dict["Duration"] = k.Duration
			dict["CreateTime"] = k.CreateTime
			dict["Status"] = k.Status
			ret = append(ret, dict)
		}
	}
	return ret
}

// 删除任务
func DelTask(taskMap *TaskMap, name string) error {
	gm := getMap(taskMap, name)
	if gm != nil {
		gm.Chan <- Stop
		close(gm.Chan)
		delMap(taskMap, name)
		return nil
	}
	return errors.New("Error: Task Name Error")
}

// 暂停任务
func PauseTask(taskMap *TaskMap, name string) error {
	gm := getMap(taskMap, name)
	if gm != nil {
		gm.Chan <- Pause
		return nil
	}
	return errors.New("Error: Task Name Error")
}

// 启动任务
func RunTask(taskMap *TaskMap, name string) error {
	gm := getMap(taskMap, name)
	if gm != nil {
		gm.Chan <- Running
		gm.Status = Running
		return nil
	}
	return errors.New("Error: Task Name Error")
}