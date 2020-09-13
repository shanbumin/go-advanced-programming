package main

import "fmt"



//todo 注意这里就不一定能打印出[你好,世界],因为通道中可以预存1个元素，所以发送不会等待接收
func main() {
	//带缓冲通道
	done := make(chan int,1)

	go func(){
		fmt.Println("你好, 世界")
		<-done  //接收
	}()

	done <- 1 //发送
}

