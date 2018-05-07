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
	"github.com/ivpusic/grpool"

	"KafkaToDruid/utils"
	"KafkaToDruid/g"
)

// goroutine pool
// 100  个 worker
// 1000 大小的队列
<<<<<<< HEAD
var Pool = grpool.NewPool(2400, 100)
=======
var Pool = grpool.NewPool(2400, 1000)
>>>>>>> c18c868e2b0c667ebc4219834ae14b58e06676b6

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

	c, err := sarama.NewConsumer(strings.Split(g.BrokerList, ","), nil)
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
		ret, err := utils.KafkaToJson(string(msg.Value), g.UriMap)
		if err != nil {
			//fmt.Println(err)
			continue
		} else {
			// 如果 druid 集群可以接收 HTTP POST 方式发送的数据，使用 DoPost()
			//druid.DoPost("miui_fast", ret)

			// 如果 druid 集群只能从 Kafka 消费数据，使用 DoProducer()
			//DoProducer(ret)

			// 任务入队
			Pool.JobQueue <- func() {
				producer("miui_fast_data_nginx_log", ret)
			}
		}
	}

	if err := c.Close(); err != nil {
		fmt.Println("Failed to close consumer: ", err)
	}
}

func conTest(topic, offset string, partition int)  {
	fmt.Println("topic:", topic, "partition", partition)
}

func DoConsumer(state string) {
	InitProducer()
	fmt.Println(g.BrokerList)
	if state == "test" {
		// 测试处理 单 topic 单 partition
		consumer("nginx_cash", "newest", 0)
	} else if state == "init" {
		// 消费十秒
		for i := 0; i < len(g.Data.Topics); i++ {
			consumer(g.Data.Topics[i].TopicName, "newest", 0)
		}
	} else {
		// 所有 topic 并遍历所有 partition
		for i := 0; i < len(g.Data.Topics); i++ {
			for j := 0; j < g.Data.Topics[i].Partition; j++ {
				go consumer(g.Data.Topics[i].TopicName, "newest", j)
			}
		}
	}
	//defer Pool.Release()
}
