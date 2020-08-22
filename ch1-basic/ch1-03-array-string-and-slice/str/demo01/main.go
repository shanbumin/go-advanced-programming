package main

import "fmt"

//在 Go 语言中，字符串的内容是不能修改的，也就是说，你不能用 s[0] 这种方式修改字符串中的 UTF-8 编码，
//如果你一定要修改，那么你可以将字符串的内容复制到一个可写的缓冲区中，然后再进行修改。
//这样的缓冲区一般是 []byte 或 []rune。
//如果要对字符串中的字节进行修改，则转换为 []byte 格式，如果要对字符串中的字符进行修改，则转换为 []rune 格式，转换过程会自动复制数据。

func main() {

	//1.修改字符串中的字节（用 []byte）
	s := "Hello 世界！"

	b := []byte(s) // 转换为 []byte，自动复制数据

	b[5] = ',' // 修改 []byte

	fmt.Printf("%s\n", s) // s 不能被修改，内容保持不变

	fmt.Printf("%s\n", b) // 修改后的数据

	fmt.Println("-------------------------")
	//2.修改字符串中的字符（用 []rune）：

	r := []rune(s) // 转换为 []rune，自动复制数据

	r[6] = '中' // 修改 []rune

	r[7] = '国' // 修改 []rune

	fmt.Println(s) // s 不能被修改，内容保持不变
	fmt.Println(string(r)) // 转换为字符串，又一次复制数据


	
}