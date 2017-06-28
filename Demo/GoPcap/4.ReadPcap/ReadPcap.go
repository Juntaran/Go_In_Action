/**
  * Author: Juntaran
  * Email:  Jacinthmail@gmail.com
  * Date:   6/28/17 7:44 PM
  */

package main

import (
	"github.com/google/gopacket/pcap"
	"log"
	"github.com/google/gopacket"
	"fmt"
)

var pcapFile 	string = "test.pcap"

func main() {
	// Open file instead of device
	handle, err := pcap.OpenOffline(pcapFile)
	defer handle.Close()
	if err != nil {
		log.Fatal(err)
	}

	// Loop through packets in file
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		fmt.Println(packet)
	}
}
