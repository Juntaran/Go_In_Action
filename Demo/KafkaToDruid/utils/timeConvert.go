/** 
  * Author: Juntaran 
  * Email:  Jacinthmail@gmail.com 
  * Date:   2018/3/27 19:40
  */

package utils

import (
	"time"
	//"strconv"
)

// nginx 时间转化为 druid 时间格式
func NginxTimeConvert(timeNg string) string {
	tm, _ := time.Parse("02/Jan/2006:15:04:05 +0800", timeNg)
	//return strconv.Itoa(int(tm.Unix()) - 8*3600)
	return time.Unix(tm.Unix()-3600*16, 0).Format("2006-01-02T15:04:05Z")
}