/**
  * Author: Juntaran
  * Email:  Jacinthmail@gmail.com
  * Date:   2017/5/11 14:29
  */

package lib

import "time"

// 原生请求结构
type RawReq struct {
	ID		int64
	Req 	[]byte
}

// 原生响应结构
type RawResp struct {
	ID  	int64
	Resp 	[]byte
	Err 	error
	Elapse 	time.Duration		// 记录耗时
}

// 结果代码类型
type RetCode int

// 保留1-1000给载荷承受方使用
const (
	RET_CODE_SUCCESS              RetCode = 0    	// 成功
	RET_CODE_WARNING_CALL_TIMEOUT         = 1001 	// 调用超时警告
	RET_CODE_ERROR_CALL                   = 2001 	// 调用错误
	RET_CODE_ERROR_RESPONSE               = 2002 	// 响应内容错误
	RET_CODE_ERROR_CALEE                  = 2003 	// 被调用方（被测软件）的内部错误
	RET_CODE_FATAL_CALL                   = 3001 	// 调用过程中发生了致命错误
)

// 调用结果的结构
type CallResult struct {
	ID 		int64				// ID
	Req 	RawReq				// 原生请求
	Resp 	RawResp				// 原生响应
	Code 	RetCode				// 响应代码
	Msg		string				// 结果成因简述
	Elapse 	time.Duration		// 耗时
}

// 载荷发生器状态的常量
const (
	STATUS_ORIGINAL uint32 = 0		// 原始
	STATUS_STARTING uint32 = 1		// 正在启动
	STATUS_STARTED  uint32 = 2		// 已经启动
	STATUS_STOPPING uint32 = 3		// 正在停止
	STATUS_STOPPED  uint32 = 4		// 已经停止
)