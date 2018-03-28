/** 
  * Author: Juntaran 
  * Email:  Jacinthmail@gmail.com 
  * Date:   2018/3/26 14:47
  */

package utils

import (
	"strings"
	"KafkaToDruid/druid"
	"fmt"
)

func KafkaToJson(testStr string) {
	s := strings.Split(testStr, "\t")

	// 校验 testStr
	if len(s) != 14 {
		fmt.Println("Error: 日志格式不对")
		return
	}
	if s[12] == "-" {
		fmt.Println("Error: upstream_response_time 为空")
		return
	}

	// 正式代码
	timestamp := NginxTimeConvert(s[5])
	var ret = "{\"time\": " + timestamp + ", \"http_host\": \"" + s[0] + "\", \"server_addr\": \"" + s[1] + "\", \"hostname\": \"" + s[2] + "\", \"remote_addr\": \"" + s[3] + "\", \"http_x_forwarded_for\": \"" + s[4] + "\", \"time_local\": \"" + s[5] + "\", \"request_uri\": \"" + s[6] + "\", \"request_length\": " + s[7] + ", \"bytes_sent\": " + s[8] + ", \"request_time\": " + s[9] + ", \"status\": " + s[10] + ", \"upstream_addr\": \"" + s[11] + "\", \"upstream_response_time\": " + s[12] + "\", \"scheme\": \"" + s[13] + "\"}"
	//fmt.Println(ret)

	// 测试无需转换格式
	// ret := testStr

	druid.DoPost("test", ret)
}
