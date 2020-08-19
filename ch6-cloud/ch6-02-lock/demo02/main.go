package main

import (
	"sync"
)

// 全局变量
var counter int

func main() {
	var wg sync.WaitGroup
	var l sync.Mutex
	//并发同时去给counter加1
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			l.Lock()
			counter++
			l.Unlock()
		}()
	}

	wg.Wait()
	//多次运行会得到不同的结果：
	println(counter)
}