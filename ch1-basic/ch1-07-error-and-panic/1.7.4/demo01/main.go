package main

import "log"



//TODO: 在非`defer`语句中执行`recover`调用是初学者常犯的错误:
func main() {
	if r := recover(); r != nil {
		log.Fatal(r)
	}

	panic(123)

	if r := recover(); r != nil {
		log.Fatal(r)
	}
}



//上面程序中两个`recover`调用都不能捕获任何异常。
//在第一个`recover`调用执行时，函数必然是在正常的非异常执行流程中，这时候`recover`调用将返回`nil`。
//发生异常时，第二个`recover`调用将没有机会被执行到，因为`panic`调用会导致函数马上执行已经注册`defer`的函数后返回。