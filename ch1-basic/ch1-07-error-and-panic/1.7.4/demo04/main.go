package main

//如果我们直接在`defer`语句中调用`MyRecover`函数又可以正常工作了：


func MyRecover() interface{} {
	return recover()
}

func main() {
	// 可以正常捕获异常
	defer MyRecover()
	panic(1)
}



