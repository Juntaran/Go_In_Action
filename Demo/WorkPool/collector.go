/**
 * Author: Juntaran
 * Email:  Jacinthmail@gmail.com
 * Date:   2017/6/11 6:10
 */

package WorkPool

import (
	"net/http"
	"time"
	"fmt"
)

// 一个有缓冲通道，可以向其发送work请求
var WorkQueue = make(chan WorkRequest, 100)

func Collector(w http.ResponseWriter, r *http.Request)  {
	// 参数判断
	if r.Method != "POST" {
		w.Header().Set("Allow", "Post")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// 获取延迟
	delay, err := time.ParseDuration(r.FormValue("delay"))
	if err != nil {
		http.Error(w, "Bad delay value: " + err.Error(), http.StatusBadRequest)
		return
	}

	// 延迟确认处于1-10秒
	if delay.Seconds() < 1 || delay.Seconds() > 10 {
		http.Error(w, "The delay must be between 1 and 10 seconds, inclusively.", http.StatusBadRequest)
		return
	}

	// 从请求取出name
	name := r.FormValue("name")
	if name == "" {
		http.Error(w, "You must specify a name.", http.StatusBadRequest)
		return
	}

	// 基于name和delay创建一个WorkRequest
	doFunc := func(config interface{}) error {
		fmt.Sprintf("Doing it")
		return nil
	}

	work := WorkRequest{
		Execute: 	doFunc,
	}

	// 把work入队
	WorkQueue <- work
	fmt.Println("Work request queued.")

	w.WriteHeader(http.StatusCreated)
	return
}