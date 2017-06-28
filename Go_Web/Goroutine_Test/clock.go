package main

import (
	"net"
	"io"
	"time"
	"log"
)

// 一个顺序执行的时钟服务器，每隔１秒把当前时间写到客户端
func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return                      // 客户端断开连接
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConn(conn)         // 如果不加go　则只有一个客户端能够收到
	}
}