package main

//递归调用

//Go语言中，函数还可以直接或间接地调用自己，也就是支持递归调用。
//Go语言函数的递归调用深度逻辑上没有限制，函数调用的栈是不会出现溢出错误的，因为Go语言运行时会根据需要动态地调整函数栈的大小。
//每个goroutine刚启动时只会分配很小的栈（4或8KB，具体依赖实现）
//根据需要动态调整栈的大小，栈最大可以达到GB级（依赖具体实现，在目前的实现中，32位体系结构为250MB,64位体系结构为1GB）。
//在Go1.4以前，Go的动态栈采用的是分段式的动态栈，通俗地说就是采用一个链表来实现动态栈，每个链表的节点内存位置不会发生变化。
//但是链表实现的动态栈对某些导致跨越链表不同节点的热点调用的性能影响较大，因为相邻的链表节点它们在内存位置一般不是相邻的，这会增加CPU高速缓存命中失败的几率。
//为了解决热点调用的CPU缓存命中率问题，Go1.4之后改用连续的动态栈实现，也就是采用一个类似动态数组的结构来表示栈。
//不过连续动态栈也带来了新的问题：当连续栈动态增长时，需要将之前的数据移动到新的内存空间，这会导致之前栈中全部变量的地址发生变化。
//虽然Go语言运行时会自动更新引用了地址变化的栈变量的指针，但最重要的一点是要明白Go语言中指针不再是固定不变的了
//todo （因此不能随意将指针保持到数值变量中，Go语言的地址也不能随意保存到不在GC控制的环境中，因此使用CGO时不能在C语言中长期持有Go语言对象的地址）。

//todo 不需要关心堆栈问题
//因为，Go语言函数的栈会自动调整大小，所以普通Go程序员已经很少需要关心栈的运行机制的。
//在Go语言规范中甚至故意没有讲到栈和堆的概念。
//我们无法知道函数参数或局部变量到底是保存在栈中还是堆中，我们只需要知道它们能够正常工作就可以了。看看下面这个例子：




//第一个函数直接返回了函数参数变量的地址——这似乎是不可以的，因为如果参数变量在栈上的话，函数返回之后栈变量就失效了，返回的地址自然也应该失效了。
//但是Go语言的编译器和运行时比我们聪明的多，它会保证指针指向的变量在合适的地方。
func f(x int) *int {
	return &x
}

//第二个函数，内部虽然调用`new`函数创建了`*int`类型的指针对象，但是依然不知道它具体保存在哪里。
//对于有C/C++编程经验的程序员需要强调的是：不用关心Go语言中函数栈和堆的问题，编译器和运行时会帮我们搞定；
//同样不要假设变量在内存中的位置是固定不变的，指针随时可能会变化，特别是在你不期望它变化的时候。

func g() int {
	x := new(int)
	return *x
}



func main() {
	
}
