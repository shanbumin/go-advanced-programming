package main

import (
	"fmt"
	"log"
)

//其实`recover`函数调用有着更严格的要求：我们必须在`defer`函数中直接调用`recover`。
//如果`defer`中调用的是`recover`函数的包装函数的话，异常的捕获工作将失败！
//比如，有时候我们可能希望包装自己的`MyRecover`函数，在内部增加必要的日志信息然后再调用`recover`，这是错误的做法：

func main() {
	defer func() {
		// 无法捕获异常
		if r := MyRecover(); r != nil {
			fmt.Println(r)
		}
	}()
	panic(1)
}

func MyRecover() interface{} {
	log.Println("trace...")
	return recover()
}