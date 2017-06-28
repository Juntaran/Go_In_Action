/**
  * Author: Juntaran
  * Email:  Jacinthmail@gmail.com
  * Date:   6/28/17 7:14 PM
  */

package __EthInfo

import (
	"github.com/google/gopacket/pcap"
	"log"
	"fmt"
)

func main() {
	// Find all devices
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Devices:")
	for _, device := range devices {
		fmt.Println("\nName: ", device.Name)
		fmt.Println("Description: ", device.Description)
		fmt.Println("Devices addresses: ", device.Description)
		for _, address := range device.Addresses {
			fmt.Println("- IP address: ", address.IP)
			fmt.Println("- Subnet mask: ", address.Netmask)
		}
	}
}
