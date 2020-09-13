package main

import "fmt"




//将管道的接收操作放到下面，这样可以避免同步事件受管道缓存大小的影响：
func main() {
	done := make(chan int,1)

	go func(){
		fmt.Println("你好, 世界")
		done <- 1 //发送
	}()
	<-done  //接收

}

