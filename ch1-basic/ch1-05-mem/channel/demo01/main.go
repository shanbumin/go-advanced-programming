package main


var done = make(chan bool)
var msg string

func aGoroutine() {
	msg = "你好, 世界"
	done <- true //发送
}

func main() {
	go aGoroutine()
	<-done//接收
	println(msg)
}

//todo 顺序如下:
//可保证打印出“hello, world”。该程序首先对`msg`进行写入，然后在`done`管道上发送同步信号，随后从`done`接收对应的同步信号，最后执行`println`函数。
