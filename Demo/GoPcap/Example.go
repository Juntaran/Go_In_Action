/**
  * Author: Juntaran
  * Email:  Jacinthmail@gmail.com
  * Date:   6/28/17 8:28 PM
  */


//需要安装libcap-devel包
package main

//+build linux

import (
"fmt"

"github.com/google/gopacket"
"github.com/google/gopacket/pcap"
)

func main() {
	// 指定监听的网络为eth0,每次捕获消息大小,是否已混合模式打开,
	if handle, err := pcap.OpenLive("ens33", 1600, true, pcap.BlockForever); err == nil {
		//设置过滤规则,即端口为80,如果指定协议则:tcp and port 80
		err = handle.SetBPFFilter("port 8000")
		if err != nil {
			fmt.Println(err)
			return
		}
		source := gopacket.NewPacketSource(handle, handle.LinkType())
		for v := range source.Packets() {
			//判断数据包是否是Payload如果是则打印,
			if payload := v.Layer(gopacket.LayerTypePayload); payload != nil {
				fmt.Println(string(payload.LayerContents()))
			}
		}
	}
}
