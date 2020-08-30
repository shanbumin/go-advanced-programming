package main

import (
	"sync/atomic"
	"time"
)

//sync/atomic`包对基本的数值类型及复杂对象的读写都提供了原子操作的支持。
//`atomic.Value`原子对象提供了`Load`和`Store`两个原子方法，
//分别用于加载和保存数据，返回值和参数都是`interface{}`类型，因此可以用于任意的自定义复杂类型。


type Config struct {
	 Name  string
	 Age    int
	 Sex    int
}
func loadConfig()*Config  {
	return &Config{
		Name: "sam",
		Age:  18,
		Sex:  19,
	}
}


func main() {

	var config atomic.Value // 保存当前配置信息

	// 初始化配置信息
	config.Store(loadConfig())

	// 启动一个后台线程, 加载更新后的配置信息
	go func() {
		for {
			time.Sleep(time.Second)
			config.Store(loadConfig())
		}
	}()

	// 用于处理请求的工作者线程始终采用最新的配置信息
	//for i := 0; i < 10; i++ {
	//	go func() {
	//		for r := range requests() {
	//			c := config.Load()
	//			// ...
	//		}
	//	}()
	//}
}
