package main

//如果我们直接在`defer`语句中调用`MyRecover`函数又可以正常工作了：




//但是，如果`defer`语句直接调用`recover`函数，依然不能正常捕获异常：
func main() {
	// 无法捕获异常
	defer recover()
	//defer func() {
	//	if r := recover(); r != nil {
	//		fmt.Println(r)
	//	}
	//}()


	panic(1)
}



//必须要和有异常的栈帧只隔一个栈帧，`recover`函数才能正常捕获异常。
//换言之，`recover`函数捕获的是祖父一级调用函数栈帧的异常（刚好可以跨越一层`defer`函数）！



