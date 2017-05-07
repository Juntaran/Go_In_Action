## Channel

原子函数和互斥锁会让并发程序的编写更为复杂
除了原子函数和互斥锁，还可以使用通道来保证对共享资源的安全访问并消除竞争状态
利用通道，通过发送接收共享资源，在goroutine之间同步
可以通过通道共享内置类型、命名类型、结构类型和引用类型的值或者指针

    // 无缓冲的整型通道
    unbuffered := make(chan int)
    unbuffered <- 1

    // 有缓冲的字符串通道
    buffered := make(chan string, 10)

一个发送语句将一个值从一个goroutine通过channel发送到另一个执行接收操作的goroutine。
发送和接收两个操作都是用<-运算符。
在发送语句中，<-运算符分割channel和要发送的值。
在接收语句中，<-运算符写在channel对象之前。一个不使用接收结果的接收操作也是合法的。

无缓冲通道是指在接收前没有任何能力保存任何值的通道
需要发送goroutine和接收goroutine同时准备好，才能完成发送、接收操作
否则，通道会导致先执行操作的goroutine阻塞

有缓冲的通道市一中在被接收之前能存储一个或多个值的通道
它不强制要求goroutine之间必须同时完成发送和接收
>接收阻塞：通道只有在没有需要接收的值时才会阻塞

>发送阻塞：通道只有在没有可用缓冲区容纳时才会阻塞