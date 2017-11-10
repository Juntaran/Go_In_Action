/** 
  * Author: Juntaran 
  * Email:  Jacinthmail@gmail.com 
  * Date:   2017/11/10 18:17
  */

package main

import (
	"TempAuth/utils"
	"fmt"
)

func main() {
	ts := utils.NowTimestamp()
	fmt.Println(ts)

	timenow := utils.TimestampToChina(ts, "2006-01-02 15:04:05")
	fmt.Println(timenow)

	ts2 := utils.ChinaToTimestamp(timenow, "2006-01-02 15:04:05")
	fmt.Println(ts2)
}