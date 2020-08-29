package main

//Go语言中，对于基础类型（非接口类型）不支持隐式的转换，我们无法将一个`int`类型的值直接赋值给`int64`类型的变量，也无法将`int`类型的值赋值给底层是`int`类型的新定义命名类型的变量。
//Go语言对基础类型的类型一致性要求可谓是非常的严格，但是Go语言对于接口类型的转换则非常的灵活。
//对象和接口之间的转换、接口和接口之间的转换都可能是隐式的转换。可以看下面的例子：

/*

var (
	a io.ReadCloser = (*os.File)(f) // 隐式转换, *os.File 满足 io.ReadCloser 接口
	b io.Reader     = a             // 隐式转换, io.ReadCloser 满足 io.Reader 接口
	c io.Closer     = a             // 隐式转换, io.ReadCloser 满足 io.Closer 接口
	d io.Reader     = c.(io.Reader) // 显式转换, io.Closer 不满足 io.Reader 接口
)
*/


func main() {
	
}
