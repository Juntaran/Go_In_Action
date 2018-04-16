/** 
  * Author: Juntaran 
  * Email:  Jacinthmail@gmail.com 
  * Date:   2018/3/26 16:33
  */

package main

import (
	"log"
	"flag"
	"os"
	"runtime/pprof"
	_ "net/http/pprof"
	
	"KafkaToDruid/kafka"
)

var State = "test2"
var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)

		defer pprof.StopCPUProfile()
	}

	if State == "test" {
		kafka.BrokerList = "localhost:9092"
	} else if State == "test2" {
		kafka.BrokerList = "localhost:9092"
	} else {
		kafka.BrokerList = kafka.Data.Brokers
	}
	var endChan = make(chan struct{}, 1)
	kafka.DoConsumer(State)
	<- endChan
}