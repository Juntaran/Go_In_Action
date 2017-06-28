/**
  * Author: Juntaran
  * Email:  Jacinthmail@gmail.com
  * Date:   6/28/17 7:47 PM
  */

package main

import (
	"time"
	"github.com/google/gopacket/pcap"
	"log"
	"fmt"
	"github.com/google/gopacket"
)

var (
	device			string			= "ens33"
	snapshotLen 	int32  			= 1024
	promiscuous		bool			= false		// 混杂
	timeout 		time.Duration	= 30 * time.Second

	filter 			string 			= "tcp and port 80"
)

func main() {
	// Open device
	handle, err := pcap.OpenLive(device, snapshotLen, promiscuous, timeout)
	defer handle.Close()
	if err != nil {
		log.Fatal(err)
	}

	// Set filter
	// Berkeley Packet Filter 伯克利封包过滤器
	err = handle.SetBPFFilter(filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Only Capture TCP port 80 packets.")

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		// Do something with a packet here
		fmt.Println(packet)
	}
}
