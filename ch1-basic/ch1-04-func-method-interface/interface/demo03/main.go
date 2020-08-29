package main

import (
	"os"
	"strings"
	"fmt"
)

//当然，我们也可以定义自己的打印格式来实现将每个字符转为大写字符后输出的效果。
//对于每个要打印的对象，如果满足了`fmt.Stringer`接口，则默认使用对象的`String`方法返回的结果打印：

//实现fmt.Stringer接口的结构体类型
type UpperString string

func (s UpperString) String() string {
	return strings.ToUpper(string(s))
}



func main() {
	//比如我们仍然可以打印一个对象出来，只要打印的对象实现fmt.Stringer接口
	fmt.Fprintln(os.Stdout, UpperString("hello, world"))
}



