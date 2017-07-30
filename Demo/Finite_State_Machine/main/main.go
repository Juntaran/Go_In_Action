/**
 * Author: Juntaran
 * Email:  Jacinthmail@gmail.com
 * Date:   2017/7/31 0:30
 */

package main

import (
	"fmt"
	. "Go_In_Action/Demo/Finite_State_Machine"
)

var (
	Poweroff        = FSMState("关闭")
	FirstGear       = FSMState("1档")
	SecondGear      = FSMState("2档")
	ThirdGear       = FSMState("3档")
	PowerOffEvent   = FSMEvent("按下关闭按钮")
	FirstGearEvent  = FSMEvent("按下1档按钮")
	SecondGearEvent = FSMEvent("按下2档按钮")
	ThirdGearEvent  = FSMEvent("按下3档按钮")

	PowerOffHandler = FSMHandler(func() FSMState {
		fmt.Println("电风扇已关闭")
		return Poweroff
	})
	FirstGearHandler = FSMHandler(func() FSMState {
		fmt.Println("电风扇开启1档，微风徐来！")
		return FirstGear
	})
	SecondGearHandler = FSMHandler(func() FSMState {
		fmt.Println("电风扇开启2档，凉飕飕！")
		return SecondGear
	})
	ThirdGearHandler = FSMHandler(func() FSMState {
		fmt.Println("电风扇开启3档，发型被吹乱了！")
		return ThirdGear
	})
)

// 电风扇
type ElectricFan struct {
	*FSM
}

// 实例化电风扇
func NewElectricFan(initState FSMState) *ElectricFan {
	return &ElectricFan{
		FSM: NewFSM(initState),
	}
}

// 入口函数
func main() {

	efan := NewElectricFan(Poweroff) // 初始状态是关闭的
	// 关闭状态
	efan.AddHandler(Poweroff, PowerOffEvent, PowerOffHandler)
	efan.AddHandler(Poweroff, FirstGearEvent, FirstGearHandler)
	efan.AddHandler(Poweroff, SecondGearEvent, SecondGearHandler)
	efan.AddHandler(Poweroff, ThirdGearEvent, ThirdGearHandler)
	// 1档状态
	efan.AddHandler(FirstGear, PowerOffEvent, PowerOffHandler)
	efan.AddHandler(FirstGear, FirstGearEvent, FirstGearHandler)
	efan.AddHandler(FirstGear, SecondGearEvent, SecondGearHandler)
	efan.AddHandler(FirstGear, ThirdGearEvent, ThirdGearHandler)
	// 2档状态
	efan.AddHandler(SecondGear, PowerOffEvent, PowerOffHandler)
	efan.AddHandler(SecondGear, FirstGearEvent, FirstGearHandler)
	efan.AddHandler(SecondGear, SecondGearEvent, SecondGearHandler)
	efan.AddHandler(SecondGear, ThirdGearEvent, ThirdGearHandler)
	// 3档状态
	efan.AddHandler(ThirdGear, PowerOffEvent, PowerOffHandler)
	efan.AddHandler(ThirdGear, FirstGearEvent, FirstGearHandler)
	efan.AddHandler(ThirdGear, SecondGearEvent, SecondGearHandler)
	efan.AddHandler(ThirdGear, ThirdGearEvent, ThirdGearHandler)

	// 开始测试状态变化
	efan.Call(ThirdGearEvent)  // 按下3档按钮
	efan.Call(FirstGearEvent)  // 按下1档按钮
	efan.Call(PowerOffEvent)   // 按下关闭按钮
	efan.Call(SecondGearEvent) // 按下2档按钮
	efan.Call(PowerOffEvent)   // 按下关闭按钮
}