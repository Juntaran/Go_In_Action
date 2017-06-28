/**
  * Author: Juntaran
  * Email:  Jacinthmail@gmail.com
  * Date:   6/28/17 7:28 PM
  */

package main

import (
	"github.com/google/gopacket/pcap"
	"time"
	"os"
	"github.com/google/gopacket/pcapgo"
	"github.com/google/gopacket/layers"
	"fmt"
	"github.com/google/gopacket"
)

var (
	device			string			= "ens33"
	snapshotLen 	uint32  		= 1024
	promiscuous		bool			= false		// 混杂
	timeout 		time.Duration	= -1 * time.Second
	packetCount		int 			= 0
)

func main() {
	// Open output pcap file and write header
	f, _ := os.Create("test.pcap")
	defer f.Close()
	w := pcapgo.NewWriter(f)
	w.WriteFileHeader(snapshotLen, layers.LinkTypeEthernet)

	// Open the device for capturing
	handle, err := pcap.OpenLive(device, int32(snapshotLen), promiscuous, timeout)
	defer handle.Close()
	if err != nil {
		fmt.Printf("Error opening device %s: %v", device, err)
		os.Exit(1)
	}

	// Start processing packets
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		// Process packet here
		fmt.Println(packet)
		w.WritePacket(packet.Metadata().CaptureInfo, packet.Data())
		packetCount ++

		// Only capture 100 and then stop
		if packetCount > 100 {
			break
		}
	}
}