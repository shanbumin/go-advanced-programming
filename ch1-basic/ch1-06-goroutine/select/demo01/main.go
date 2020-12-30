package main

import (
	"fmt"
	"time"
)

//Go语言中不同Goroutine之间主要依靠管道进行通信和同步。
//要同时处理多个管道的发送或接收操作，我们需要使用`select`关键字（这个关键字和网络编程中的`select`函数的行为类似）。
//当`select`有多个分支时，会随机选择一个可用的管道分支，如果没有可用的管道分支则选择`default`分支，否则会一直保存阻塞状态。


//基于`select`实现的管道的超时判断：
func main() {

	in:=make(chan int)
	go func() {
		for i:=0;i<100;i++{
			time.Sleep(500 * time.Millisecond)
			in<- i
		}
	}()

	//for{
		select {
		case v := <-in:
			fmt.Println(v)
		case <-time.After(time.Second):
			fmt.Println("timeout")
			return // 超时
		}
	//}




}
