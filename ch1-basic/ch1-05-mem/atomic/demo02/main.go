package main

import (
	"sync"
	"sync/atomic"
)

var total uint64

//用互斥锁来保护一个数值型的共享资源，麻烦且效率低下。
//标准库的`sync/atomic`包对原子操作提供了丰富的支持。我们可以重新实现上面的例子：
func worker(wg *sync.WaitGroup) {
	defer wg.Done()

	var i uint64
	for i = 0; i <= 100; i++ {
		//`atomic.AddUint64`函数调用保证了`total`的读取、更新和保存是一个原子操作，因此在多线程中访问也是安全的。
		atomic.AddUint64(&total, i)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go worker(&wg)
	go worker(&wg)
	wg.Wait()
}

