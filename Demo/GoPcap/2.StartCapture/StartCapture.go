/**
  * Author: Juntaran
  * Email:  Jacinthmail@gmail.com
  * Date:   6/28/17 7:21 PM
  */

package main

import (
	"time"
	"github.com/google/gopacket/pcap"
	"log"
	"github.com/google/gopacket"
	"fmt"
)

var (
	device			string			= "ens33"
	snapshot_len	int32			= 1024
	promiscuous		bool			= false		// 混杂
	timeout 		time.Duration	= 30 * time.Second
)

func main() {
	// Open device
	handle, err := pcap.OpenLive(device, snapshot_len, promiscuous, timeout)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	// Use the handle as a packet source to process all packets
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		// Process packet here
		fmt.Println(packet)
	}
}