package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)


//在C语言中，`printf`只能将几种有限的基础数据类型打印到文件对象中。
//但是Go语言灵活接口特性，`fmt.Fprintf`却可以向任何自定义的输出流对象打印，可以打印到文件或标准输出、也可以打印到网络、甚至可以打印到一个压缩文件；
//同时，打印的数据也不仅仅局限于语言内置的基础类型，任意隐式满足`fmt.Stringer`接口的对象都可以打印，不满足`fmt.Stringer`接口的依然可以通过反射的技术打印。



//我们可以通过定制自己的输出对象，将每个字符转为大写字符后输出：
type UpperWriter struct {
	io.Writer
}
func (p *UpperWriter) Write(data []byte) (n int, err error) {
	return p.Writer.Write(bytes.ToUpper(data))
}




func main() {
	//只要实现io.Writer接口都可以作为第一个参数(输出对象)
	fmt.Fprintln(&UpperWriter{os.Stdout}, "hello, world")
}



