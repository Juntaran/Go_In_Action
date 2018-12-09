# Tcp_Forward

TCP端口转发  

## 测试方法：

1. 启动端口转发程序，监听8000端口
   
       ./Tcp_Forward -l=:8000 -d=127.0.0.1:1788,127.0.0.1:1789
    
    
2. 使用[WorkPool](https://github.com/Juntaran/github.com/Juntaran/Go_In_Action/tree/master/Demo/WorkPool)进行测试

        ./queued -n=4 -http=localhost:8000          // 被监听端口
        ./queued -n=4 -http=localhost:1788          // 转发端口1
        ./queued -n=4 -http=localhost:1789          // 转发端口2
   
3. 使用`curl`或`Postman`发起request请求

        curl -v -X POST "localhost:8000/work?delay=1s&name=foo"
        
    > 会依次转发到不同的端口，先是1789，再1788，再1789...
    
## Reference:
* [Mefellows](http://blog.csdn.net/fyxichen/article/details/51505542)
