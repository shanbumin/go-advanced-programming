package main

import (
	"log"
	"sync"
	"time"
)

//我们通过`close`来关闭`cancel`管道向多个Goroutine广播退出的指令。
//不过这个程序依然不够稳健：当每个Goroutine收到退出指令退出时一般会进行一定的清理工作，但是退出的清理工作并不能保证被完成，
//因为`main`线程并没有等待各个工作Goroutine退出工作完成的机制。我们可以结合`sync.WaitGroup`来改进:



func worker(wg *sync.WaitGroup, cancel chan bool) {
	defer func() {
		//todo 清理工作
		log.Println("该G做完了清理工作了...")
	}()
	defer wg.Done()

	for {
		select {
		default:
			log.Println("hello")
		case <-cancel:
			return
		}
	}

}

//现在每个工作者并发体的创建、运行、暂停和退出都是在`main`函数的安全控制之下了。
func main() {
	cancel := make(chan bool)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(&wg, cancel)
	}

	time.Sleep(time.Second)
	close(cancel)
	wg.Wait()
}
