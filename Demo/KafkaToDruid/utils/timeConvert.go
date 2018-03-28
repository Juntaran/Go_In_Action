/** 
  * Author: Juntaran 
  * Email:  Jacinthmail@gmail.com 
  * Date:   2018/3/27 19:40
  */

package utils

import (
	"time"
	"strconv"
)

// nginx 时间转化为时间戳
func NginxTimeConvert(timeNg string) string {
	tm, _ := time.Parse("02/Jan/2006:15:04:05 +0800", timeNg)
	return strconv.Itoa(int(tm.Unix()) - 8*3600)
}