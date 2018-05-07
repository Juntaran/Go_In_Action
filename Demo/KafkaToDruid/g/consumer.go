/** 
  * Author: Juntaran 
  * Email:  Jacinthmail@gmail.com 
  * Date:   2018/4/19 17:48
  */

package g

import (
	"log"
	"os"
	"strconv"
	"strings"
	"os/signal"
	"fmt"

	"github.com/Shopify/sarama"

	"KafkaToDruid/utils"
)

var Exit chan struct{}

// 消费者
func consumer(topic, offset string, partition int) {
	fmt.Println("start init map consumer")
	brokerlist := BrokerList

	var logger = log.New(os.Stderr, "", log.LstdFlags)

	var (
		initialOffset int64
		offsetError   error
	)
	switch offset {
	case "oldest":
		initialOffset = sarama.OffsetOldest
	case "newest":
		initialOffset = sarama.OffsetNewest
	default:
		initialOffset, offsetError = strconv.ParseInt(offset, 10, 64)
	}

	if offsetError != nil {
		logger.Fatalln("Invalid initial offset:", offset)
	}

	c, err := sarama.NewConsumer(strings.Split(brokerlist, ","), nil)
	if err != nil {
		logger.Fatalln(err)
	}

	pc, err := c.ConsumePartition(topic, int32(partition), initialOffset)
	if err != nil {
		logger.Fatalln(err)
	}

	go func() {
		signals := make(chan os.Signal, 1)
		signal.Notify(signals, os.Kill, os.Interrupt)
		<-signals
		pc.AsyncClose()
	}()

	for msg := range pc.Messages() {
		s := strings.Split(string(msg.Value), "\t")[6]
		ret := utils.DoUri(s, 2)

		if v, ok := TempMap.Get(ret); ok {
			TempMap.Set(ret, v.(int)+1)
		} else {
			TempMap.Set(ret, 1)
		}
	}

	if err := c.Close(); err != nil {
		fmt.Println("Failed to close consumer: ", err)
	}
}
