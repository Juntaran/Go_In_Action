# 从 kafka 读取数据

对每个 topic 的 每个 partition 开启一个 goroutine 拉取，并转换为 json 格式