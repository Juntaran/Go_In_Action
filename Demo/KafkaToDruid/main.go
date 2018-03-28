/** 
  * Author: Juntaran 
  * Email:  Jacinthmail@gmail.com 
  * Date:   2018/3/26 16:33
  */

package main

import (
	"KafkaToDruid/kafka"
)

var State = "test"

func main() {
	if State == "test" {
		kafka.BrokerList = "localhost:9092"
	} else {
		kafka.BrokerList = kafka.Data.Brokers
	}
	var endChan = make(chan struct{}, 1)
	kafka.DoKafka(State)
	<- endChan
}