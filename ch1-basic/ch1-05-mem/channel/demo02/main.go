package main


var done = make(chan bool)
var msg string

func aGoroutine() {
	msg = "你好, 世界"
	close(done)
}

func main() {
	go aGoroutine()
	<-done
	println(msg)
}
//若在关闭Channel后继续从中接收数据，接收者就会收到该Channel返回的零值。
//因此在这个例子中，用`close(c)`关闭管道代替`done <- false`依然能保证该程序产生相同的行为。
