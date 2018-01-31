/** 
  * Author: Juntaran 
  * Email:  Jacinthmail@gmail.com 
  * Date:   2018/1/31 11:10
  */

package main

import (
	"log"
	"time"
	"Go_In_Action/Demo/BlockchainBasic"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	go func() {
		genesisBlock := BlockchainBasic.Block{
			Index: 			0,
			Timestamp: 		time.Now().String(),
			BPM: 			0,
			Hash: 			"",
			PrevHash: 		"",
		}
		spew.Dump(genesisBlock)
		BlockchainBasic.Blockchain = append(BlockchainBasic.Blockchain, genesisBlock)
	}()
	log.Fatal(BlockchainBasic.Run())
}