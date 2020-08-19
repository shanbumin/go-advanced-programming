package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
)

func main() {

	//数组不仅仅可以用于数值类型，还可以定义字符串数组、结构体数组、函数数组、接口数组、管道数组等等：

	//1.字符串数组
	/*
	var s1 = [2]string{"hello", "world"}
	var s2 = [...]string{"你好", "世界"}
	var s3 = [...]string{1: "世界", 0: "你好", }
    */


	//2.结构体数组
	/*
	var line1 [2]image.Point
	var line2 = [...]image.Point{image.Point{X: 0, Y: 0}, image.Point{X: 1, Y: 1}}
	var line3 = [...]image.Point{{0, 0}, {1, 1}}
	*/


	//3.图像解码器数组
	/*
	var decoder1 [2]func(io.Reader) (image.Image, error)
	var decoder2 = [...]func(io.Reader) (image.Image, error){
	   png.Decode,
	   jpeg.Decode,
	}
	 */

	//4.接口数组
	/*
	var unknown1 [2]interface{}
	var unknown2 = [...]interface{}{123, "你好"}
	*/


	//5.管道数组
	/*
	var chanList = [2]chan int{}
    */


	//6.我们还可以定义一个空的数组：
    /*
	var d [0]int       // 定义一个长度为0的数组
	var e = [0]int{}   // 定义一个长度为0的数组
	var f = [...]int{} // 定义一个长度为0的数组
     */
	//长度为0的数组在内存中并不占用空间。空数组虽然很少直接使用，但是可以用于强调某种特有类型的操作时避免分配额外的内存空间，比如用于管道的同步操作：

	   c1 := make(chan [0]int)
	   go func() {
		  fmt.Println("c1")
		  c1 <- [0]int{}
	   }()
	   <-c1

	//在这里，我们并不关心管道中传输数据的真实类型，其中管道接收和发送操作只是用于消息的同步。
	//对于这种场景，我们用空数组来作为管道类型可以减少管道元素赋值时的开销。
	//当然一般更倾向于用无类型的匿名结构体代替：

   c2 := make(chan struct{})
   go func() {
      fmt.Println("c2")
      c2 <- struct{}{} // struct{}部分是类型, {}表示对应的结构体值
   }()
   <-c2


}
