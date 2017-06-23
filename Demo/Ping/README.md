# Ping

Linux:

    go build -o myPing .
    ./myPing -host=www.baidu.com -n=3

Windows:

    cd ~\Go_In_Action\Demo\Ping\main
    go build -o myPing.exe main.go
    main.exe -host=www.baidu.com -n=3
 
当`-host`选项不填写时，默认 Ping `127.0.0.1`  
当`-n`选项不填写时，不会停止 Ping  


## Reference:
* [JieLinDee](http://blog.csdn.net/fyxichen/article/details/51995647)

