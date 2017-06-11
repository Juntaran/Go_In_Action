# Prime

从2开始每找到一个素数就标记所有能被该素数整除的所有数  
直到没有可以标记的数，剩下的都是素数  
CSP方式解决该问题  

## CSP

Go的并发属于`CSP`并发模型的一种实现  
CSP的核心概念是  
`不要通过共享内存来通信，通过通信共享内存`  
Goroutine和Channel

## Reference:
* [今日头条Go建千亿级微服务的实践](https://zhuanlan.zhihu.com/p/26695984)