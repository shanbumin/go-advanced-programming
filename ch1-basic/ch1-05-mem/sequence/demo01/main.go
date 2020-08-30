package main

//什么是顺序一致性的内存模型?
var a string
var done bool

func setup() {
	a = "hello, world"
	done = true
	//done = false
}


//我们创建了`setup`线程，用于对字符串`a`的初始化工作，初始化完成之后设置`done`标志为`true`。
//`main`函数所在的主线程中，通过`for !done {}`检测`done`变为`true`时，认为字符串初始化工作完成，然后进行字符串的打印工作。


//todo 下面的解释可能有点晦涩难懂额,在`setup`线程中必然是顺序一致性执行的，即如果done变成true了，则a的值必然已经赋值为"hello, world"
//todo 但是在`main`线程中就是无法观测到了，即哪怕这个时候done被`setup`线程设置为true了,在`main`中也可能观测到a的值为空字符串额
//但是Go语言并不保证在`main`函数中观测到的对`done`的写入操作发生在对字符串`a`的写入的操作之后，因此程序很可能打印一个空字符串。
//更糟糕的是，因为两个线程之间没有同步事件，`setup`线程对`done`的写入操作甚至无法被`main`线程看到，`main`函数有可能陷入死循环中。

func main() {
	go setup()
	//done为false的话将一直执行,死循环下去
	for !done {
		//fmt.Printf("a的值为%#v\r\n",a)
	}
	print(a)
}




