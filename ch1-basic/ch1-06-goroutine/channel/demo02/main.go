package main

import (
	"fmt"
	"sync"
)

//下面是修复后的代码：
func main() {
	var mu sync.Mutex

	mu.Lock()
	go func(){
		fmt.Println("你好, 世界")
		mu.Unlock()
	}()

	mu.Lock()
}


