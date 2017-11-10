/** 
  * Author: Juntaran 
  * Email:  Jacinthmail@gmail.com 
  * Date:   2017/11/10 17:48
  */

package utils

import "time"

// 获取当前时间戳
func NowTimestamp() int64 {
	return time.Now().Unix()
}

// 时间戳转化为当前中国时间 GMT，返回string
// 自己指定日期格式 例如 2006-01-02 15:04:05
func TimestampToChina(timestamp int64, format string) string {
	return time.Unix(timestamp, 0).Format(format)
}

// 中国时间字符串转化为时间戳
func ChinaToTimestamp(tmChina string, format string) int64 {
	tm, _ := time.Parse(format, tmChina)
	ts := tm.Unix() - 3600 * 8
	return ts
}