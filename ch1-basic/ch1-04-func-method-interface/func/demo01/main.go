package main

//在Go语言中，函数是第一类对象，我们可以将函数保持到变量中。
//函数主要有具名和匿名之分，包级函数一般都是具名函数，具名函数是匿名函数的一种特例。
//当然，Go语言中每个类型还可以有自己的方法，方法其实也是函数的一种。



// 具名函数
func Add(a, b int) int {
	return a+b
}





func main() {
	// 匿名函数
	var Add2 = func(a, b int) int {
		return a+b
	}

	_ =Add2

}
