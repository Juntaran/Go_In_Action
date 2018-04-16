# KafkaToDruid


## 使用前提:  

``` sh
go get -u github.com/Shopify/sarama
go get -u gopkg.in/yaml.v2
```

## Nginx 日志格式如下:  

``` nginx
log_format milog '$http_host\t$server_addr\t$hostname\t$remote_addr\t$http_x_forwarded_for\t$time_local\t$request_uri\t$request_length\t$bytes_sent\t$request_time\t$status\t$upstream_addr\t$upstream_response_time\t$scheme';
```

对应着:  

```
0  http_host
1  server_addr
2  hostname
3  remote_addr
4  http_x_forwarded_for
5  time_local
6  request_uri
7  request_length
8  bytes_sent
9  request_time
10 status
11 upstream_addr
12 upstream_response_time
13 scheme
```

## 日志处理流程:  

NginxLog -> td_agent -> lcs -> kafka -> `kafkaConsumer -> kafka2json -> postDruid` -> tranquility -> druid

其中高亮为 `KafkaToDruid` 中间件的文件名，相关操作可以直接查看  

## 使用方法:  

druid 部署详情参考 [github/Juntaran/Druid](https://github.com/Juntaran/Note/blob/master/Data_Mining-Machine_Learning/Druid/Druid%E9%83%A8%E7%BD%B2.md)  

``` sh
# 开启 zk
cd zookeeper-3.4.10
cp conf/zoo_sample.cfg conf/zoo.cfg
./bin/zkServer.sh start

# 开启 druid
cd druid-0.11.0
bin/init
java `cat conf-quickstart/druid/historical/jvm.config | xargs` -cp "conf-quickstart/druid/_common:conf-quickstart/druid/historical:lib/*" io.druid.cli.Main server historical
java `cat conf-quickstart/druid/broker/jvm.config | xargs` -cp "conf-quickstart/druid/_common:conf-quickstart/druid/broker:lib/*" io.druid.cli.Main server broker
java `cat conf-quickstart/druid/coordinator/jvm.config | xargs` -cp "conf-quickstart/druid/_common:conf-quickstart/druid/coordinator:lib/*" io.druid.cli.Main server coordinator
java `cat conf-quickstart/druid/overlord/jvm.config | xargs` -cp "conf-quickstart/druid/_common:conf-quickstart/druid/overlord:lib/*" io.druid.cli.Main server overlord
java `cat conf-quickstart/druid/middleManager/jvm.config | xargs` -cp "conf-quickstart/druid/_common:conf-quickstart/druid/middleManager:lib/*" io.druid.cli.Main server middleManager

# 开启 tranquility
cp $GOPATH/src/KafkaToDruid/druid/miuiServer.json tranquility-distribution-0.8.0/conf
cd tranquility-distribution-0.8.0
bin/tranquility server -configFile ./conf/miuiServer.json

# 开启 KafkaToDruid 中间件
go run main.go > ret.txt
```

## 中间件处理策略:  

1. 热更新组件，定时查询 yaml 配置文件 md5 是否有改变，如果有增加会开启新的 goroutine 支持，不能减少
2. 读取 yaml 文件，获得 `kafka-topic`、`topic` 以及每个 topic 对应的 `partition`  
3. 分别对每个 topic 的 partition 开启一个新 goroutine  
4. 每个 goroutine 连接 kafka 拉取对应 `topic` 对应 `partition` 的数据  
5. 数据转换为 json 格式，该格式与 `tranquility` 的 `miuiServer.json` 对应  
6. json 以 `HTTP POST` 方式打入 `tranquility`，由 tranquility 自动打入 druid  


## 实际线上流程:  

因为公司只支持从 kafka 集群拉取数据，所以我们经过中间件转发数据后，还需要打回 kafka  
[kafka] -> 中间件 -> [kafka_new -> traquility -> druid]  

如果支持 http post 方式打入集群，则可以通过修改 `kafka/kafkaConsumer.go` 直接略过 producer 使用 http 打入 `traquility`
