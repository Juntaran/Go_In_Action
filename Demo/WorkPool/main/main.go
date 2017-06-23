/**
 * Author: Juntaran
 * Email:  Jacinthmail@gmail.com
 * Date:   2017/6/11 6:30
 */

package main

import (
	"flag"
	"fmt"
	"Go_In_Action/Demo/WorkPool"
	"net/http"
)

var (
	NWorkers = flag.Int("n", 4, "The number of workers to start")
	HTTPAddr = flag.String("http", "localhost:8000", "Address to listen for HTTP requests on")
)

func main() {
	// 解析命令行参数
	flag.Parse()

	// 启动dispatcher
	fmt.Println("Starting the dispatcher.")
	WorkPool.StartDispatcher(*NWorkers)

	// 注册collector为一个HTTP处理函数
	fmt.Println("Registering the collector.")
	http.HandleFunc("/work", WorkPool.Collector)

	// 启动HTTP服务器
	fmt.Println("HTTP server listeing on", *HTTPAddr)
	if err := http.ListenAndServe(*HTTPAddr, nil); err != nil {
		fmt.Println(err.Error())
	}
}

// https://github.com/mefellows/golang-worker-example