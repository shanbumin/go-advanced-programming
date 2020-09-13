package main

import (
	"fmt"
	"sync"
)

//对于这种要等待N个线程完成后再进行下一步的同步操作有一个简单的做法，就是使用`sync.WaitGroup`来等待一组事件：


func main() {
	var wg sync.WaitGroup

	// 开N个后台打印线程
	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func() {
			fmt.Println("你好, 世界")
			wg.Done()
		}()
	}

	// 等待N个后台线程完成
	wg.Wait()
}

//其中`wg.Add(1)`用于增加等待事件的个数，必须确保在后台线程启动之前执行（如果放到后台线程之中执行则不能保证被正常执行到）。
//当后台线程完成打印工作之后，调用`wg.Done()`表示完成一个事件。`main`函数的`wg.Wait()`是等待全部的事件完成。


