/**
  * Author: Juntaran
  * Email:  Jacinthmail@gmail.com
  * Date:   2017/5/8 01:24
  */

package main

import (
	"net"
	"strings"
	"fmt"
	"sync"
	"time"
	"io"
	"bytes"
	"strconv"
	"math"
	"math/rand"
)

const (
	SERVER_NETWORK = "tcp"
	SERVER_ADDRESS = "127.0.0.1:8086"
	DELIMITER      = '\t'				// 数据边界
)

var wg sync.WaitGroup

// 从连接中读取一段以数据分界符为结尾的数据
func read(conn net.Conn) (string, error) {
	readBytes := make([]byte, 1)
	var buffer bytes.Buffer
	for {
		_, err := conn.Read(readBytes)
		if err != nil {
			return "", err
		}
		readByte := readBytes[0]
		if readByte == DELIMITER {
			break
		}
		buffer.WriteByte(readByte)
	}
	return buffer.String(), nil
}

// 发送数据
func write(conn net.Conn, content string) (int, error) {
	var buffer bytes.Buffer
	buffer.WriteString(content)
	buffer.WriteByte(DELIMITER)
	return conn.Write(buffer.Bytes())
}

// 处理连接
func handleConn(conn net.Conn)  {
	defer func() {
		conn.Close()
		wg.Done()
	}()
	// handleConn首先试图从连接中循环读取数据
	for {
		// 设置超时
		conn.SetDeadline(time.Now().Add(10 * time.Second))
		// 从连接中读取一段以数据分界符为结尾的数据
		strReq, err := read(conn)
		if err != nil {
			if err == io.EOF {
				printServerLog("The connection is closed by another side.")
			} else {
				printServerLog("Read Error: %s", err)
			}
			break
		}
		printServerLog("Received request: %s.", strReq)

		// 检查数据块是否可以转换为一个int32类型的值
		intReq, err := strToInt32(strReq)
		if err != nil {
			// 如果不能，向客户端发送一条错误信息
			n, err := write(conn, err.Error())
			printServerLog("Send error message (written &d bytes): %s.", n, err)
			continue
		}
		// 如果能，计算立方根
		floatResp := cbrt(intReq)
		respMsg := fmt.Sprintf("The cube root of %d is %f.", intReq, floatResp)
		n, err := write(conn, respMsg)
		if err != nil {
			printServerLog("Write Error: %s", err)
		}
		printServerLog("Sent response (written %d bytes): %s.", n, respMsg)
	}
}


// 服务端程序
func serverGo()  {
	var listener net.Listener
	listener, err := net.Listen(SERVER_NETWORK, SERVER_ADDRESS)
	if err != nil {
		printServerLog("Listen Error: %s", err)
		return
	}
	defer listener.Close()
	printServerLog("Got listener for the server. (local address: %s)", listener.Addr())

	// 循环监听，一旦成功获得监听器，就可以开始等待客户端连接请求
	for {
		// 阻塞直到有新连接进来
		conn, err := listener.Accept()
		if err != nil {
			printServerLog("Accept Error: %s", err)
		}
		printServerLog("Established a connection with a client application. (remote address: %s)", conn.RemoteAddr())
		// 启动一个新的goroutine来并发执行handleConn函数，这是很有必要的
		// 为了快速、独立地处理已经建立的每一个连接，应该并发执行
		go handleConn(conn)
	}
}

// 客户端程序
func clientGo(id int)  {
	defer wg.Done()
	conn, err := net.DialTimeout(SERVER_NETWORK, SERVER_ADDRESS, 2 * time.Second)
	// 连接不成功记录日志直接返回
	if err != nil {
		printClientLog(id, "Dial Error: %s", err)
		return
	}
	defer conn.Close()
	printClientLog(id, "Connected to server. (remote address: %s, local address: %s)", conn.RemoteAddr(), conn.LocalAddr())
	time.Sleep(200 * time.Millisecond)

	// 每个客户端发送的请求数据块数量定位5
	requestNumber := 5
	conn.SetDeadline(time.Now().Add(5 * time.Millisecond))

	// 发送5个请求数据块
	for i := 0; i < requestNumber; i++ {
		req := rand.Int31()
		n, err := write(conn, fmt.Sprintf("%d", req))
		if err != nil {
			printClientLog(id, "Write Error: %s", err)
			continue
		}
		printClientLog(id, "Sent request (written %d bytes): %d.", n, req)
	}

	// 发送完后，接收响应数据块
	for j := 0; j < requestNumber; j++ {
		strResp, err := read(conn)
		if err != nil {
			if err == io.EOF {
				printClientLog(id, "The connection is closed by another side.")
			} else {
				printClientLog(id, "Read Error: %s", err)
			}
			break
		}
		printClientLog(id, "Received response: %s.", strResp)
	}
}



// Log
func printLog(role string, sn int, format string, args ...interface{})  {
	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}
	fmt.Printf("%s[%d]: %s", role, sn, fmt.Sprintf(format, args...))
}

// 服务端日志
func printServerLog(format string, args ...interface{})  {
	printLog("Server", 0, format, args...)
}

// 客户端日志
func printClientLog(sn int, format string, args ...interface{})  {
	printLog("Client", sn, format, args...)
}

// strToInt32
func strToInt32(str string) (int32, error) {
	num, err := strconv.ParseInt(str, 10, 0)
	if err != nil {
		return 0, fmt.Errorf("\"%s\" is not integer", str)
	}
	if num > math.MaxInt32 || num < math.MinInt32 {
		return 0, fmt.Errorf("%d is not 32-bit intger", num)
	}
	return int32(num), nil
}

// 计算立方根
func cbrt(param int32) float64 {
	return math.Cbrt(float64(param))
}

func main() {
	wg.Add(3)
	go serverGo()
	time.Sleep(500 * time.Millisecond)
	go clientGo(1)
	go clientGo(2)
	wg.Wait()
}