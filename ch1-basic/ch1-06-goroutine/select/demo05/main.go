package main

import (
	"fmt"
	"time"
)

//我们通过`select`和`default`分支可以很容易实现一个Goroutine的退出控制:

func worker(cancel chan bool) {
	for {
		select {
		default:
			fmt.Println("hello")
			// 正常工作
		case <-cancel:
			// 退出
		}
	}
}

//todo: 停止单个goroutine
func main() {
	cancel := make(chan bool)
	//开启一个G，但是我们传递一个取消通道，用来在main中控制这个G的退出
	go worker(cancel)

	time.Sleep(time.Second)
	cancel <- true
}
