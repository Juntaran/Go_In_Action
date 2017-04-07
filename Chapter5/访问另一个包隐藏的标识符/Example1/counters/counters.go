/**
  * Author: Juntaran
  * Email:  Jacinthmail@gmail.com
  * Date:   2017/4/7 19:13
  */

package counters

// counters包提供了告警计数器功能

// alertCounter 是一个未公开类型
type alertCounter int

// New创建并返回一个未公开的alertCounter类型的值
// New是一个公开的函数，返回值是alterCounter类型的值
func New(value int) alertCounter {
	return alertCounter(value)
}
