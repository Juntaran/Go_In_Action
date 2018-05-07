/** 
  * Author: Juntaran 
  * Email:  Jacinthmail@gmail.com 
  * Date:   2018/4/17 18:00
  */

package g

import (
	"sort"
	"fmt"
	"time"

	"github.com/orcaman/concurrent-map"
)

var UriMap = cmap.New()		// uri 统计 map  每个 +1
var UriArr PairList

var TempMap	= cmap.New()

type Pair struct {
	Key   	string
	Value 	int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool { return p[i].Value > p[j].Value }

func cos()  {
	v, ok := UriMap.Get("aaa")
	if ok == true {
		UriMap.Set("aaa", v.(int)+1)
	} else {
		UriMap.Set("aaa", 1)
	}
}

func f(offset string) {
	for i := 0; i < 8; i++ {
		consumer("nginx_data", offset, i)
	}
	//cos()
}

// 初始化 uri 统计
func InitUriCount(top int)  {
	go f("newest")
	timeout := time.After(time.Second * 10)
	for {
		select {
		case <-timeout:
			fmt.Println("超时")
			sortMapByValue(TempMap, top)
			fmt.Println("UriArr:", UriArr)
			return
		}
	}
}

// 初始化 uri 统计
func initUriCountOld()  {
	go f("oldest")
	timeout := time.After(time.Second * 10)
	for {
		select {
		case <-timeout:
			fmt.Println("超时")
			return
		}
	}
}

// 对 map 排序，并取 top 写入 UriArr，同时更新 TempMap
func sortMapByValue(m cmap.ConcurrentMap, top int) {
	var p PairList

	for item := range m.IterBuffered() {
		p = append(p, Pair{item.Key, item.Val.(int)})
	}

	if len(p) == 0 {
		fmt.Println("读取 kafka newest 为空，选择 oldest")
		//fmt.Println("当前 UriArr:", UriArr)
		//fmt.Println("当前 UriMapKeys:", UriMap.Keys())
		initUriCountOld()
		for item := range m.IterBuffered() {
			p = append(p, Pair{item.Key, item.Val.(int)})
		}
	}

	sort.Sort(p)
	if len(UriArr) > top {
		UriArr = p[:top]
	} else {
		UriArr = p
	}

	// 更新 UriMap，只保留 top 个
	if len(UriArr) > 0 {
		//fmt.Println("开始替换 UriMap")
		//fmt.Println("UriArr:", UriArr)
		var tempMap = cmap.New()
		for _, v := range UriArr {
			tempMap.Set(v.Key, v.Value)
		}
		//fmt.Println("tempMap:", tempMap)
		//fmt.Println("替换前 UriMap:", UriMap.Keys())
		UriMap = tempMap
		//fmt.Println("替换后 UriMap:", UriMap.Keys())
	}

	// 清理 TempMap
	//fmt.Println("清理前 TempMap:", TempMap.Keys())
	//fmt.Println("清理 TempMap")
	TempMap = cmap.New()
	//fmt.Println("清理后 TempMap:", TempMap.Keys())
}

// 定时根据 UriMap 更新 UriArr
func UpdateUri(sec, top int)  {
	InitUriCount(top)
	// 每隔 sec 执行一次 initUriCount()
	ticker := time.NewTicker(time.Second * time.Duration(sec))
	go func() {
		for _ = range ticker.C {
			go InitUriCount(top)
		}
	}()
	ch := make(chan int)
	<- ch
}