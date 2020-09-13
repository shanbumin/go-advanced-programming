package main

import "fmt"

//使用`sync.Mutex`互斥锁同步是比较低级的做法。我们现在改用无缓存的管道来实现同步：


func main() {
	done := make(chan int)

	go func(){
		fmt.Println("你好, 世界")
		<-done  //接收
	}()

	done <- 1 //发送
}


//根据Go语言内存模型规范，对于从无缓冲Channel进行的接收，发生在对该Channel进行的发送完成之前。
//因此，后台线程`<-done`接收操作完成之后，`main`线程的`done <- 1`发送操作才可能完成（从而退出main、退出程序），而此时打印工作已经完成了。