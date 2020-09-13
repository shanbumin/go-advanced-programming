package main

import (
	"chai2010.cn/gobook/ch1-basic/ch1-06-goroutine/publish-subscribe/pubsub"
	"fmt"
	"strings"
	"time"
)

//下面的例子中，有两个订阅者分别订阅了全部主题和含有"golang"的主题：
func main() {
	//1.构建一个发布者对象, 可以设置发布超时时间和缓存队列的长度
	p := pubsub.NewPublisher(100*time.Millisecond, 10)
	defer p.Close()

	//2.添加一个新的订阅者，订阅全部主题
	all := p.Subscribe()
	//3.添加一个新的订阅者，订阅过滤器筛选后的主题
	golang := p.SubscribeTopic(func(v interface{}) bool {
		if s, ok := v.(string); ok {
			return strings.Contains(s, "golang")
		}
		return false
	})



    //4.发布两条订阅信息
    //todo 发布的消息是否会进入上述的两个管道中，会先通过过滤器筛选
	p.Publish("hello,  world!")
	p.Publish("hello, golang!")



	//--------------模拟两个订阅者------------------------------------------------------
	go func() {
		for  msg := range all {
			fmt.Println("all:", msg)
		}
	} ()

	go func() {
		for  msg := range golang {
			fmt.Println("golang:", msg)
		}
	} ()

	// 运行一定时间后退出
	time.Sleep(3 * time.Second)
}


