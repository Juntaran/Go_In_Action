/** 
  * Author: Juntaran 
  * Email:  Jacinthmail@gmail.com 
  * Date:   2018/2/2 01:27
  */

package main

import (
	"flag"
	"strings"
	"fmt"
	"Go_In_Action/Demo/Blockchain"
	"net/http"
	"log"
)

func main() {
	serverPort := flag.String("port", "8000", "http port number where server will run")
	flag.Parse()

	blockchain := Blockchain.NewBlockchain()
	nodeID := strings.Replace(Blockchain.PseudoUUID(), "-", "", -1)

	log.Printf("Starting gochain HTTP Server. Listening at port %q", *serverPort)

	http.Handle("/", Blockchain.NewHandler(blockchain, nodeID))
	http.ListenAndServe(fmt.Sprintf(":%s", *serverPort), nil)
}
