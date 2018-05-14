/** 
  * Author: Juntaran 
  * Email:  Jacinthmail@gmail.com 
  * Date:   2018/5/11 20:08
  */

package main

import (
	"strings"
	"strconv"
	"fmt"
	"sort"
)

func transIpToNum(ip string) int {
	s := strings.Split(ip, ".")
	a0, _ := strconv.Atoi(s[0])
	a1, _ := strconv.Atoi(s[1])
	a2, _ := strconv.Atoi(s[2])
	a3, _ := strconv.Atoi(s[3])
	var ret int
	ret += a0 << 24
	ret += a1 << 16
	ret += a2 << 8
	ret += a3
	return ret
}

func transNumToIp(num int) string {
	//fmt.Println(num)
	//fmt.Printf("%b\n", num)
	a1 := num - num >> 8 << 8
	//fmt.Printf("num >> 8 << 8: %b\n", num >> 8 << 8)
	//fmt.Printf("a1: %b\n", a1)
	a2 := num - num >> 16 << 16 - a1
	//fmt.Printf("num >> 16 << 16: %b\n", num >> 16 << 16)
	//fmt.Printf("a2: %b\n", a2)
	a3 := num - num >> 24 << 24 - a1 - a2
	//fmt.Printf("num >> 24 << 24: %b\n", num >> 24 << 24)
	//fmt.Printf("a3: %b\n", a3)
	a4 := num - num >> 32 << 32 - a1 - a2 - a3
	//fmt.Printf("num >> 32 << 32: %b\n", num >> 32 << 32)
	//fmt.Printf("a4: %b\n", a4)

	ret := fmt.Sprintf("%d.%d.%d.%d", a4>>24, a3>>16, a2>>8, a1)
	return ret
}


func sortIp(ipList []string) []string {
	var ips = make([]int, len(ipList))
	for k, v := range ipList {
		ips[k] = transIpToNum(v)
	}
	//fmt.Println(ips)
	sort.Ints(ips)
	//fmt.Println(ips)

	var ret = make([]string, len(ipList))
	for k, v := range ips {
		ret[k] = transNumToIp(v)
	}
	return ret
}

func main() {
	ipList := []string{"1.0.0.1", "1.0.0.20", "192.168.213.124", "1.0.0.3"}
	fmt.Println("排序前:")
	for _, v := range ipList {
		fmt.Println(v)
	}
	ret := sortIp(ipList)
	fmt.Println("\n排序后:")
	for _, v := range ret {
		fmt.Println(v)
	}
}