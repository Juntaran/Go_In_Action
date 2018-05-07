/** 
  * Author: Juntaran 
  * Email:  Jacinthmail@gmail.com 
  * Date:   2018/4/10 17:55
  */

package kafka

import (
	"fmt"
	"os"
	"strings"
	"io/ioutil"
	"github.com/Shopify/sarama"
	"github.com/rcrowley/go-metrics"

	"KafkaToDruid/g"
)

var ProducerG 	sarama.SyncProducer
var ConfigG		= sarama.NewConfig()

func InitProducer() {
	//config := sarama.NewConfig()
	var err error

	ConfigG.Producer.RequiredAcks = sarama.WaitForAll
	ConfigG.Producer.Return.Successes = true
	ConfigG.Producer.Partitioner = sarama.NewRandomPartitioner

	ProducerG, err = sarama.NewSyncProducer(strings.Split(g.BrokerList, ","), ConfigG)

	// 初始化全局变量
	//ProducerG = producer

	if err != nil {
		fmt.Println("Failed to open Kafka producer:", err)
	} else {
		fmt.Println("Init Producer Success")
	}
	//defer func() {
	//	if err := producer.Close(); err != nil {
	//		logger.Println("Failed to close Kafka producer cleanly:", err)
	//	}
	//}()
}

func producer(topic, value string) {
	var (
		//key         			= ""
		//partitioner 			= "random"
		partition 		int32  	= -1
		//logger 					= log.New(os.Stderr, "", log.LstdFlags)
		showMetrics 			= false 	// Output metrics on successful publish to stderr
		silent      			= false 	// Turn off printing the message's topic, partition, and offset to stdout
	)

	//config := sarama.NewConfig()
	//config.Producer.RequiredAcks = sarama.WaitForAll
	//config.Producer.Return.Successes = true
	//
	//config.Producer.Partitioner = sarama.NewRandomPartitioner

	//switch partitioner {
	//case "":
	//	if partition >= 0 {
	//		config.Producer.Partitioner = sarama.NewManualPartitioner
	//	} else {
	//		config.Producer.Partitioner = sarama.NewHashPartitioner
	//	}
	//case "hash":
	//	config.Producer.Partitioner = sarama.NewHashPartitioner
	//case "random":
	//	config.Producer.Partitioner = sarama.NewRandomPartitioner
	//case "manual":
	//	config.Producer.Partitioner = sarama.NewManualPartitioner
	//	if partition == -1 {
	//		fmt.Println("-partition is required when partitioning manually")
	//		//printUsageErrorAndExit("-partition is required when partitioning manually")
	//	}
	//default:
	//	fmt.Printf("Partitioner %s not supported.\n", partitioner)
	//	//printUsageErrorAndExit(fmt.Sprintf("Partitioner %s not supported.", partitioner))
	//}


	message := &sarama.ProducerMessage{Topic: topic, Partition: int32(partition)}

	//if key != "" {
	//	message.Key = sarama.StringEncoder(key)
	//}

	// 构建 message
	if value != "" {
		message.Value = sarama.StringEncoder(value)
	} else if stdinAvailable() {
		bytes, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			fmt.Println("Failed to read data from the standard input:", err)
		}
		message.Value = sarama.ByteEncoder(bytes)
	} else {
		fmt.Println("-value is required, or you have to provide the value on stdin")
		//printUsageErrorAndExit("-value is required, or you have to provide the value on stdin")
	}

	//producer, err := sarama.NewSyncProducer(strings.Split(BrokerList, ","), config)
	//if err != nil {
	//	fmt.Println("Failed to open Kafka producer:", err)
	//}
	//defer func() {
	//	if err := producer.Close(); err != nil {
	//		logger.Println("Failed to close Kafka producer cleanly:", err)
	//	}
	//}()

	// 发送数据
	//partition, offset, err := producer.SendMessage(message)
	partition, _, err := ProducerG.SendMessage(message)
	if err != nil {
		fmt.Println("Failed to produce message:", err)
	} else if !silent {
		//fmt.Printf("topic=%s\tpartition=%d\toffset=%d\n", topic, partition, offset)
	}
	if showMetrics {
		metrics.WriteOnce(ConfigG.MetricRegistry, os.Stderr)
	}
}

//func printUsageErrorAndExit(message string) {
//	fmt.Fprintln(os.Stderr, "ERROR:", message)
//	fmt.Fprintln(os.Stderr)
//	fmt.Fprintln(os.Stderr, "Available command line options:")
//	flag.PrintDefaults()
//	os.Exit(64)
//}

func stdinAvailable() bool {
	stat, _ := os.Stdin.Stat()
	return (stat.Mode() & os.ModeCharDevice) == 0
}

func proTest()  {
	for i := 0; i < len(g.Data.Topics); i++ {
		fmt.Println("miui_fast_" + strings.Split(g.Data.Topics[i].TopicName, "_")[1] + "_nginx_log")
	}
}

// 不能为每一条数据开启一个 goroutine
// 应该做成 goroutine pool
func DoProducer(jsonStr string)  {
	//fmt.Println(BrokerList)
	for i := 0; i < len(g.Data.Topics); i++ {
		//producer("fast_" + strings.Split(Data.Topics[i].TopicName, "_")[1], jsonStr)
		producer("miui_fast_cash_nginx_log", jsonStr)
	}
}