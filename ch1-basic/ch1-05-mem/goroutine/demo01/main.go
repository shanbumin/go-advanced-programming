package main


//`go`语句会在当前Goroutine对应函数返回前创建新的Goroutine. 例如:

var a string

func f() {
   print(a)
}

func hello() {
   a = "hello, world"
   go f()
}

//执行`go f()`语句创建Goroutine和`hello`函数是在同一个Goroutine中执行,
//根据语句的书写顺序可以确定Goroutine的创建发生在`hello`函数返回之前,

//但是新创建Goroutine对应的`f()`的执行事件和`hello`函数返回的事件则是不可排序的，也就是并发的。
//调用`hello`可能会在将来的某一时刻打印`"hello, world"`，也很可能是在`hello`函数执行完成后才打印。


func main() {
	hello()
}
