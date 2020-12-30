package main

import "fmt"

//同样，如果是在嵌套的`defer`函数中调用`recover`也将导致无法捕获异常：


func main() {
	defer func() {
		defer func() {
			// 无法捕获异常
			if r := recover(); r != nil {
				fmt.Println(r)
			}
		}()
	}()
	panic(1)
}
//2层嵌套的`defer`函数中直接调用`recover`和1层`defer`函数中调用包装的`MyRecover`函数一样，
//都是经过了2个函数帧才到达真正的`recover`函数，
//这个时候Goroutine的对应上一级栈帧中已经没有异常信息。


