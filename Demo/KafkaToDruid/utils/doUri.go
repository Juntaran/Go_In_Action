/**
  * Author: Juntaran
  * Email:  Jacinthmail@gmail.com
  * Date:   2018/4/11 11:11
  */

package utils

// 拆分 uri location
func DoUri(uri string, splitN int) string {
	if uri == "" {
		return uri
	}
	count 	:= 0
	pos   	:= 0
	lastPos := 0
	for k, v := range uri {
		pos = k
		if v == '?' {
			return uri[:pos]
		}
		if v == '/' {
			lastPos = k
			count ++
		}
		if v == '.' {
			if lastPos == 0 {
				return "/"
			}
			return uri[:lastPos]
		}
		if count == splitN + 1 {
			return uri[:pos]
		}
	}
	return uri[:pos+1]
}
