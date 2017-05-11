/**
  * Author: Juntaran
  * Email:  Jacinthmail@gmail.com
  * Date:   2017/5/11 15:06
  */

package lib

// Goroutine票池的接口
type GoTickets interface {
	// 拿走一张票
	Take()
	// 归还一张票
	Return()
	// 票池是否已被激活
	Active() bool
	// 票的总数
	Total() uint32
	// 剩余的票数
	Remainder() uint32
}
