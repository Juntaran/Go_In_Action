# KafkaToDruid


## 使用前提:  

``` sh
go get -u github.com/Shopify/sarama
go get -u gopkg.in/yaml.v2
```

## Nginx 日志格式如下:  

``` nginx
log_format testlog '$http_host\t$server_addr\t$hostname\t$remote_addr\t$http_x_forwarded_for\t$time_local\t$request_uri\t$request_length\t$bytes_sent\t$request_time\t$status\t$upstream_addr\t$upstream_response_time\t$scheme';
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
cp $GOPATH/src/KafkaToDruid/druid/test.json tranquility-distribution-0.8.0/conf
cd tranquility-distribution-0.8.0
bin/tranquility server -configFile ./conf/test.json

# 开启 KafkaToDruid 中间件
go run main.go > ret.txt
```

## 中间件处理策略:  

1. 读取 yaml 文件，获得 `kafka-topic`、`topic` 以及每个 topic 对应的 `partition`  
2. 分别对每个 topic 的 partition 开启一个新 goroutine  
3. 每个 goroutine 连接 kafka 拉取对应 `topic` 对应 `partition` 的数据  
4. 数据转换为 json 格式，该格式与 `tranquility` 的 `miuiServer.json` 对应  
5. json 以 `HTTP POST` 方式打入 `tranquility`，由 tranquility 自动打入 druid  


## 存在的问题:  

kafka 的数据不会实时打入，会有几分钟到十几分钟的延迟  
本地搭建的 kafka 0.8.11 没有此问题，线上 kafka 存在该问题  