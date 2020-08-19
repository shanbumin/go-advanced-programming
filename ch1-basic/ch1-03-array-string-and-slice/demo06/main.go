package main

import "fmt"

func main() {

	//字符串虽然不是切片，但是支持切片操作，不同位置的切片底层也访问的同一块内存数据（因为字符串是只读的，相同的字符串面值常量通常是对应同一个字符串常量）：

	s := "hello, world"
	hello := s[:5]
	world := s[7:]

	fmt.Println(hello)
	fmt.Println(world)

	fmt.Println("--------------")

	s1 := "hello, world"[:5]
	s2 := "hello, world"[7:]
	fmt.Println(s1)
	fmt.Println(s2)



	
}
