# WorkPool

## Golang工作池  

可以修改`collector`里的具体处理函数  

Linux:

    go build -o queued .
    ./queued -n=4 -http=localhost:8000

Windows:

    cd ~\github.com/Juntaran/Go_In_Action\Demo\WorkPool\main
    go build -o queued.exe main.go
    queued.exe -n=4 -http=localhost:8000
    

样例程序启动后会在使用8000端口制造一个小型HTTP服务端，  
可以使用`curl`来对其进行测试：

    curl -v -X POST "localhost:8000/work?delay=1s&name=foo"

也可以使用`Postman`等工具

## Reference:
* [Mefellows](https://github.com/mefellows/golang-worker-example)
