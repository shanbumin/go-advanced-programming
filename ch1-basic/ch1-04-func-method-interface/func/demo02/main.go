package main

import "fmt"

//Go语言中的函数可以有多个参数和多个返回值，参数和返回值都是以【传值】的方式和被调用者交换数据。
//在语法上，函数还支持可变数量的参数，可变数量的参数必须是最后出现的参数，可变数量的参数其实是一个切片类型的参数。




//1.多个参数和多个返回值
func Swap(a, b int) (int, int) {
	return b, a
}

//2.可变数量的参数
func Sum(a int, more ...int) int {
	//more 对应 []int 切片类型
	for _, v := range more {
		a += v
	}
	return a
}

//3.当可变参数是一个空接口类型时，调用者是否解包可变参数会导致不同的结果：
func Print(a ...interface{}) {
	fmt.Println(a...)
}

//4.不仅函数的参数可以有名字，也可以给函数的返回值命名：
func Find(m map[int]int, key int) (value int, ok bool) {
   value, ok = m[key]
   return
}





func main() {
	var a = []interface{}{123, "abc"}

	//第一个`Print`调用时传入的参数是`a...`，等价于直接调用`Print(123, "abc")`。
	//第二个`Print`调用传入的是未解包的`a`，等价于直接调用`Print([]interface{}{123, "abc"})`。
	Print(a...) // 123 abc
	Print(a)    // [123 abc]


}
