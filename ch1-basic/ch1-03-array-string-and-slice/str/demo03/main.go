package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {


	s := "hello, world"
	s1 := "hello, world"[:5]
	s2 := "hello, world"[7:]
	//字符串和数组类似，内置的`len`函数返回字符串的长度。也可以通过`reflect.StringHeader`结构访问字符串的长度（这里只是为了演示字符串的结构，并不是推荐的做法）：


	// StringHeader是字符串的运行时表示形式。
	fmt.Printf("%#v\r\n",(*reflect.StringHeader)(unsafe.Pointer(&s)))
	fmt.Println("len(s):", (*reflect.StringHeader)(unsafe.Pointer(&s)).Len)   // 12
	fmt.Println("len(s1):", (*reflect.StringHeader)(unsafe.Pointer(&s1)).Len) // 5
	fmt.Println("len(s2):", (*reflect.StringHeader)(unsafe.Pointer(&s2)).Len) // 5





	
}
