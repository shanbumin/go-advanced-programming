package main

import "fmt"



//当有多个管道均可操作时，`select`会随机选择一个管道。基于该特性我们可以用`select`实现一个生成随机数序列的程序：
func main() {
	ch := make(chan int)
	go func() {
		for {
			select {
			case ch <- 0:
			case ch <- 1:
			}
		}
	}()

	for v := range ch {
		fmt.Println(v)
	}
}



