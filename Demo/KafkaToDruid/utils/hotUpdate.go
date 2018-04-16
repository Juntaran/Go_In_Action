/** 
  * Author: Juntaran 
  * Email:  Jacinthmail@gmail.com 
  * Date:   2018/3/29 10:25
  */

package utils

import (
	"crypto/md5"
	"os"
	"fmt"
	"bufio"
	"io"
	"encoding/hex"
	"time"
)

var ConfigMd5 string

// 计算文件 md5
func calMd5(filePath string) string {
	fmt.Println("ConfigMd5", ConfigMd5)
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error: Open Config File Error", err)
		return ""
	}
	r := bufio.NewReader(f)
	h := md5.New()
	_, err = io.Copy(h, r)
	f.Close()
	if err != nil {
		fmt.Println("Error: io Copy", err)
		return ""
	}
	return hex.EncodeToString(h.Sum(nil))
}

// 热更新组件，只能增加 topic 不能减少、不能修改已有 partition
func HotUpdate()  {
	// 初始化 ConfigMd5
	ConfigMd5 = calMd5("kafka/kafka.yml")
	// 每 1 分钟检查一次配置文件是否有变化
	ticker := time.NewTicker(time.Second * 5)
	var configMd5 string
	go func() {
		for _ = range ticker.C {
			fmt.Println("ticked at", time.Now())
			configMd5 = calMd5("kafka/kafka.yml")
			if configMd5 != ConfigMd5 {
				fmt.Println("配置文件发生变更")
				ConfigMd5 = configMd5
				// do sth
			}
		}
	}()
}

func main() {
	HotUpdate()
	var endChan = make(chan struct{}, 1)
	<- endChan
}