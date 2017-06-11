/**
 * Author: Juntaran
 * Email:  Jacinthmail@gmail.com
 * Date:   2017/6/11 13:05
 */

package main

import (
	"net"
	"fmt"
	"log"
	"sync"
	"io"
	"flag"
	"strings"
	"os"
)

var locker sync.Mutex
var trueList []string
var ip string
var list string

func main() {
	flag.StringVar(&ip, "l", ":8000", "-l=0.0.0.0:8000 指定服务监听的端口")
	flag.StringVar(&list, "d", "127.0.0.1:1789,127.0.0.1:1788", "-d=127.0.0.1:1789,127.0.0.1:1788 指定后端的IP和端口,多个用','隔开")
	flag.Parse()
	trueList = strings.Split(list, ",")
	if len(trueList) <= 0 {
		fmt.Println("后端IP和端口不能空,或者无效")
		os.Exit(1)
	}
	Tcp_Forward_Server()
}

func Tcp_Forward_Server()  {
	// (bind) -> listen -> accept -> (read/write) -> close
	listen, err := net.Listen("tcp", ip)
	if err != nil {
		log.Println("Listen Error:", err)
		return
	}
	log.Println("Listen Success")
	defer listen.Close()

	// 循环accept
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Println("Accept Error:", err)
			continue
		}
		log.Println("Accept Success")
		fmt.Println(conn.RemoteAddr(), conn.LocalAddr())
		go handle(conn)
	}
}

func handle(sconn net.Conn) {
	defer sconn.Close()
	ip, ok := getIP()
	if !ok {
		return
	}
	dconn, err := net.Dial("tcp", ip)
	if err != nil {
		log.Printf("Connect %v Fail: %v", ip, err)
		return
	}
	ExitChan := make(chan bool, 1)
	go func(sconn net.Conn, dconn net.Conn, Exit chan bool) {
		// 把sconn复制给dconn
		_, err := io.Copy(dconn, sconn)
		if err != nil {
			fmt.Printf("Copy Data from %v Fail: %v", ip, err)
			ExitChan <- true
		}
	} (sconn, dconn, ExitChan)
	<-ExitChan
	dconn.Close()
}

func getIP() (string, bool) {
	locker.Lock()
	defer locker.Unlock()

	if len(trueList) < 1 {
		return "", false
	}
	ip := trueList[0]
	trueList = append(trueList[1:], ip)
	return ip, true
}