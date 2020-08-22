package main

import (
	"fmt"
	"unicode/utf8"
)

//todo 下面分别用伪代码简单模拟Go语言对字符串内置的一些操作，这样对每个操作的处理的时间复杂度和空间复杂度都会有较明确的认识。





func main() {

	s:="世"
	r, size := utf8.DecodeRuneInString(s) //Unicode编码10进制,占的字节数
	fmt.Println(r,size)
	fmt.Println("----------------------------")

	//utf8.EncodeRune(buf, r)

}
