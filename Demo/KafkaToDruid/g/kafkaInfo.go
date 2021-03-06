/** 
  * Author: Juntaran 
  * Email:  Jacinthmail@gmail.com 
  * Date:   2018/3/23 14:59
  */

package g

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

var Data = GetKafkaData()
var BrokerList string

type KafkaTopic struct {
	Brokers 	string
	Topics		[]Topic
}

type Topic struct {
	TopicName	string
	Partition 	int
}

func GetKafkaData() KafkaTopic {
	data, _ := ioutil.ReadFile("kafka/kafka.yml")
	//fmt.Println(string(data))

	t := KafkaTopic{}
	yaml.Unmarshal(data, &t)
	//fmt.Println("初始数据:", t)

	return t

	//d, _ := yaml.Marshal(&t)
	//fmt.Println(string(d))
}

