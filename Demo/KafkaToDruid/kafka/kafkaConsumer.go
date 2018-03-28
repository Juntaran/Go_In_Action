/** 
  * Author: Juntaran 
  * Email:  Jacinthmail@gmail.com 
  * Date:   2018/3/22 17:11
  */

package kafka

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"strings"

	"github.com/Shopify/sarama"

	"KafkaToDruid/utils"
)

var Data = utils.GetKafkaData()
var BrokerList string

// 消费者
func consumer(topic, offset string, partition int) {

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

	c, err := sarama.NewConsumer(strings.Split(BrokerList, ","), nil)
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
		//fmt.Printf("Partation: %d\n", partition)
		//fmt.Printf("Offset:    %d\n", msg.Offset)
		//fmt.Printf("Key:       %s\n", string(msg.Key))
		//fmt.Printf("Value:     %s\n", string(msg.Value))
		//fmt.Println()
		utils.KafkaToJson(string(msg.Value))
	}

	if err := c.Close(); err != nil {
		fmt.Println("Failed to close consumer: ", err)
	}
}

func conTest(topic, offset string, partition int)  {
	fmt.Println("topic:", topic, "partition", partition)
}

func DoKafka(state string) {
	fmt.Println(BrokerList)
	if state == "test" {
		// 测试处理 单 topic 单 partition
		consumer("test", "newest", 0)
	} else {
		// 所有 topic 并遍历所有 partition
		for i := 0; i < len(Data.Topics); i++ {
			for j := 0; j < Data.Topics[i].Partition; j++ {
				go consumer(Data.Topics[i].TopicName, "newest", j)
			}
		}
	}
}