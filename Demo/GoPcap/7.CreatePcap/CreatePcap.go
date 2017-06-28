/**
  * Author: Juntaran
  * Email:  Jacinthmail@gmail.com
  * Date:   6/28/17 8:03 PM
  */

package main

import (
	"time"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"log"
	"github.com/google/gopacket/layers"
	"net"
)

var (
	device			string			= "ens33"
	snapshotLen 	int32  			= 1024
	promiscuous		bool			= false		// 混杂
	timeout 		time.Duration	= 30 * time.Second
	buffer       	gopacket.SerializeBuffer
	options      	gopacket.SerializeOptions
)

func main() {
	// Open device
	handle, err := pcap.OpenLive(device, snapshotLen, promiscuous, timeout)
	defer handle.Close()
	if err != nil {
		log.Fatal(err)
	}

	// Send raw bytes over wire
	rawBytes := []byte{10, 20, 30}
	err2 := handle.WritePacketData(rawBytes)
	if err2 != nil {
		log.Fatal(err2)
	}

	// 创建一个格式正确的包, 填写MAC地址、IP地址等, 其余为空
	buffer = gopacket.NewSerializeBuffer()
	gopacket.SerializeLayers(buffer, options,
		&layers.Ethernet{},
		&layers.IPv4{},
		&layers.TCP{},
		gopacket.Payload(rawBytes),
	)
	outgoingPacket := buffer.Bytes()

	// Send packet
	err3 := handle.WritePacketData(outgoingPacket)
	if err3 != nil {
		log.Fatal(err3)
	}

	// 填写信息
	ethernetLayer := &layers.Ethernet{
		SrcMAC: net.HardwareAddr{0xFF, 0xAA, 0xFA, 0xAA, 0xFF, 0xAA},
		DstMAC: net.HardwareAddr{0xBD, 0xBD, 0xBD, 0xBD, 0xBD, 0xBD},
	}

	ipLayer := &layers.IPv4{
		SrcIP: 	net.IP{127, 0, 0, 1},
		DstIP:	net.IP{114, 114, 114, 114},
	}

	tcpLayer := &layers.TCP{
		SrcPort: layers.TCPPort(4321),
		DstPort: layers.TCPPort(80),
	}

	// And create the packet with the layers
	buffer = gopacket.NewSerializeBuffer()
	gopacket.SerializeLayers(
		buffer,
		options,
		ethernetLayer,
		ipLayer,
		tcpLayer,
		gopacket.Payload(rawBytes),
	)
	outgoingPacket = buffer.Bytes()
}