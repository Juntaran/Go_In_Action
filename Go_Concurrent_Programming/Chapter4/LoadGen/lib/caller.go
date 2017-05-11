/**
  * Author: Juntaran
  * Email:  Jacinthmail@gmail.com
  * Date:   2017/5/11 15:03
  */

package lib

import "time"

// 调用器接口
type Caller interface {
	BuildReq() RawReq
	Call(req []byte, timeoutNS time.Duration) ([]byte, error)
	CheckResp(rawReq RawReq, rawResp RawResp) *CallResult
}