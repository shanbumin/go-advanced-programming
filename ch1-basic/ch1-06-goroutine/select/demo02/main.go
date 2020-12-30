package main

import (
	"fmt"
	"time"
)




//通过`select`的`default`分支实现非阻塞的管道发送或接收操作：
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
	 default:
		 // 没有数据
	 }
 //}



}
