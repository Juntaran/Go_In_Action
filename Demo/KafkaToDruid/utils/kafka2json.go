/** 
  * Author: Juntaran 
  * Email:  Jacinthmail@gmail.com 
  * Date:   2018/3/26 14:47
  */

package utils

import (
	"strings"
	"errors"
	"bytes"
	"strconv"

	"github.com/orcaman/concurrent-map"
)

func KafkaToJson(testStr string, m cmap.ConcurrentMap) (string, error) {
	// testStr = "api.ad.xiaomi.com	10.132.10.54	c4-miui-l7-data03.bj	180.106.200.165	-	25/Mar/2018:13:38:45 +0800	/desktopfolder/uploadLogAdSDK	2168	253	0.001	200	10.132.28.50:8098	-0.001	Dalvik/2.1.0 (Linux; U; Android 7.1.2; MI 5X MIUI/V9.2.1.0.NDBCNEK)	http"

	s := strings.Split(testStr, "\t")

	// 校验 testStr
	if len(s) != 15 {
		return "", errors.New("Error: 日志格式不对")
	}
	if s[12] == "-" {
		s[12] = "0"
		return "", errors.New("Error: 日志格式不对")
	}
	if _, ok := m.Get(s[6]); !ok {
		return "", errors.New("Error:" + s[6] + "不在白名单")
	}

	// 正式代码
	timestamp := NginxTimeConvert(s[5])
	//var ret = "{\"time\": \"" + timestamp + "\", \"http_host\": \"" + s[0] + "\", \"server_addr\": \"" + s[1] + "\", \"hostname\": \"" + s[2] + "\", \"remote_addr\": \"" + s[3] + "\", \"http_x_forwarded_for\": \"" + s[4] + "\", \"time_local\": \"" + s[5] + "\", \"request_uri\": \"" + DoUri(s[6], 2) + "\", \"request_length\": " + s[7] + ", \"bytes_sent\": " + s[8] + ", \"request_time\": " + s[9] + ", \"status\": " + s[10] + ", \"upstream_addr\": \"" + s[11] + "\", \"upstream_response_time\": " + s[12] + ", \"http_user_agent\": \"" + s[13] + "\", \"scheme\": \"" + s[14] + "\"}"

	// 测试无需转换格式
	// ret := testStr

	var buf = bytes.Buffer{}
	buf.WriteString("{\"time\": \"")
	buf.WriteString(timestamp)
	buf.WriteString("\", \"http_host\": \"")
	buf.WriteString(s[0])
	buf.WriteString("\", \"server_addr\": \"")
	buf.WriteString(s[1])
	buf.WriteString("\", \"hostname\": \"")
	buf.WriteString(s[2])
	buf.WriteString("\", \"remote_addr\": \"")
	buf.WriteString(s[3])
	buf.WriteString("\", \"http_x_forwarded_for\": \"")
	buf.WriteString(s[4])
	buf.WriteString("\", \"time_local\": \"")
	buf.WriteString(s[5])
	buf.WriteString("\", \"request_uri\": \"")
	buf.WriteString(DoUri(s[6], 2))
	buf.WriteString("\", \"request_length\": ")
	buf.WriteString(s[7])
	buf.WriteString(", \"bytes_sent\": ")
	buf.WriteString(s[8])
	buf.WriteString(", \"request_time\": ")
	rt, _ := strconv.ParseFloat(s[9], 32)
	rTime := strconv.FormatFloat(rt*1000, 'E', -1, 64)[:5]
	if strings.Contains(rTime, "E") {
		buf.WriteString("0.000")
	} else {
		buf.WriteString(rTime)
	}
	buf.WriteString(", \"status\": ")
	buf.WriteString(s[10])
	buf.WriteString(", \"upstream_addr\": \"")
	buf.WriteString(s[11])
	buf.WriteString("\", \"upstream_response_time\": ")
	buf.WriteString(s[12])
	buf.WriteString(", \"http_user_agent\": \"")
	buf.WriteString(s[13])
	buf.WriteString("\", \"scheme\": \"")
	buf.WriteString(s[14])
	buf.WriteString("\"}")

	var ret = buf.String()

	//fmt.Println("正常日志:", s)
	//fmt.Println("经转换后:", ret)

	return ret, nil
}
