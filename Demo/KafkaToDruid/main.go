/** 
  * Author: Juntaran 
  * Email:  Jacinthmail@gmail.com 
  * Date:   2018/3/26 16:33
  */

package main

import (
	"KafkaToDruid/kafka"
	_ "net/http/pprof"
	"flag"
	"os"
	"runtime/pprof"
	"log"
	"time"
	"KafkaToDruid/g"
)

// 原数据: api.ad.xiaomi.com	10.118.20.52	c3-miui-l7-data07.bj	123.103.40.5	-	25/Mar/2018:13:39:42 +0800	/getAds	2908	11062	0.119	200	10.118.28.30:8086	0.119	Apache-HttpClient/4.5 (Java/1.8.0_73)	http
// 转换后: {"http_host": "api.ad.xiaomi.com", "server_addr": "10.118.20.52", "hostname": "c3-miui-l7-data07.bj", "remote_addr"`: "123.103.40.5", "http_x_forwarded_for": "-", "time_local": "25/Mar/2018:13:39:42 +0800", "request_uri": "/getAds", "request_length": 2908, "bytes_sent": 11062, "request_time": 0.119, "status": 200, "upstream_addr": "10.118.28.30:8086", "upstream_response_time": 0.119, "http_user_agent": "Apache-HttpClient/4.5 (Java/1.8.0_73)", "scheme": "http"}

var State = "run"
var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")


func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)

		defer pprof.StopCPUProfile()
	}

	//go func() {
	//	// 定时器每分钟计算处理的 producer
	//	ticker := time.NewTicker(time.Second)
	//	for _ = range ticker.C {
	//		fmt.Println("Count:", kafka.Count)
	//	}
	//}()

	if State == "test" {
		g.BrokerList = "localhost:9092"
	} else if State == "test2" {
		g.BrokerList = "localhost:9092"
	} else {
		g.BrokerList = g.Data.Brokers
	}

	// 先 init map，1小时更新一次
	go g.UpdateUri(60, 20)

	// 一分钟后开始消费
	time.Sleep(time.Minute)
	var endChan = make(chan struct{}, 1)
	kafka.DoConsumer(State)
	<- endChan
}