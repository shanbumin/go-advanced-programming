package main

import (
	"fmt"
	"time"
)

// 生产者: 生成 factor 整数倍的序列
func Producer(factor int, out chan<- int) {
	for i := 0; ; i++ {
		out <- i*factor
	}
}

// 消费者
func Consumer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

//我们开启了2个`Producer`生产流水线，分别用于生成3和5的倍数的序列。
//然后开启1个`Consumer`消费者线程，打印获取的结果。
//我们通过在`main`函数休眠一定的时间来让生产者和消费者工作一定时间。
//正如前面一节说的，这种靠休眠方式是无法保证稳定的输出结果的。

func main() {
	ch := make(chan int, 64) // 成果队列

	go Producer(3, ch) // 生成 3 的倍数的序列
	go Producer(5, ch) // 生成 5 的倍数的序列
	go Consumer(ch)    // 消费 生成的队列

	// 运行一定时间后退出
	time.Sleep(5 * time.Second)
}


