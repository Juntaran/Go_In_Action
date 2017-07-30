/**
 * Author: Juntaran
 * Email:  Jacinthmail@gmail.com
 * Date:   2017/7/30 22:52
 */

package Finite_State_Machine

import (
	"sync"
	"fmt"
)

type FSMState string				// 状态
type FSMEvent string 				// 事件
type FSMHandler func() FSMState 	// 处理方法，返回状态

// 有限状态机
type FSM struct {
	lock 		sync.Mutex			// 互斥锁
	state 		FSMState			// 当前状态
	handlers	map[FSMState]map[FSMEvent]FSMHandler	// 处理map集，每一个状态可以触发有限个事件
}

// 获取当前状态
func (fsm *FSM) getState() FSMState {
	return fsm.state
}

// 设置当前状态
func (fsm *FSM) setState(newState FSMState) {
	fsm.state = newState
}

// 某状态添加事件处理方法
func (fsm *FSM) AddHandler(state FSMState, event FSMEvent, handler FSMHandler) *FSM {
	if _, ok := fsm.handlers[state]; !ok {
		fsm.handlers[state] = make(map[FSMEvent]FSMHandler)
	}
	if _, ok := fsm.handlers[state][event]; ok {
		fmt.Printf("[警告] 状态(%s)事件(%s)已定义过", state, event)
	}
	fsm.handlers[state][event] = handler
	return fsm
}

// 事件处理
func (fsm *FSM) Call(event FSMEvent) FSMState {
	fsm.lock.Lock()
	defer fsm.lock.Unlock()
	events := fsm.handlers[fsm.getState()]
	if events == nil {
		return fsm.state
	}
	if fn, ok := events[event]; ok {
		oldState := fsm.getState()
		fsm.setState(fn())
		newState := fsm.getState()
		fmt.Println("状态从 [", oldState, "] 变成 [", newState, "]")
	}
	return fsm.getState()
}

// 实例化FSM
func NewFSM(initState FSMState) *FSM {
	return &FSM{
		state: 		initState,
		handlers:	make(map[FSMState]map[FSMEvent]FSMHandler),
	}
}