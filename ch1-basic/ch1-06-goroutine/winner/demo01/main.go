package main

import (
	"fmt"
	"time"
)

//假设我们想快速地搜索“golang”相关的主题，我们可能会同时打开Bing、Google或百度等多个检索引擎。
//当某个搜索最先返回结果后，就可以关闭其它搜索页面了。
//因为受网络环境和搜索引擎算法的影响，某些搜索引擎可能很快返回搜索结果，某些搜索引擎也可能等到他们公司倒闭也没有完成搜索。
//我们可以采用类似的策略来编写这个程序：

func searchByBing(string)string  {
	time.Sleep(10 * time.Second)
	return  "bing"
}

func searchByGoogle(string)string  {
	time.Sleep(20 * time.Second)
	return  "google"
}

func searchByBaidu(string)string  {
	time.Sleep(1 * time.Second)
	return  "baidu"
}


func main() {

	ch := make(chan string, 32)

	go func() {
		ch <- searchByBing("golang")
	}()
	go func() {
		ch <- searchByGoogle("golang")
	}()
	go func() {
		ch <- searchByBaidu("golang")
	}()

	fmt.Println(<-ch)
}


//首先，我们创建了一个带缓存的管道，管道的缓存数目要足够大，保证不会因为缓存的容量引起不必要的阻塞。
//然后我们开启了多个后台线程，分别向不同的搜索引擎提交搜索请求。
//当任意一个搜索引擎最先有结果之后，都会马上将结果发到管道中（因为管道带了足够的缓存，这个过程不会阻塞）。但是最终我们只从管道取第一个结果，也就是最先返回的结果。
//通过适当开启一些冗余的线程，尝试用不同途径去解决同样的问题，最终以赢者为王的方式提升了程序的相应性能。



