package main

import "image/color"

//Go语言不支持传统面向对象中的继承特性，而是以自己特有的组合方式支持了方法的继承。Go语言中，通过在结构体内置匿名的成员来实现继承：


type Point struct{
	X, Y float64
}

//通过嵌入匿名的成员，我们不仅可以继承匿名成员的内部成员，而且可以继承匿名成员类型所对应的方法。
//我们一般会将Point看作基类，把ColoredPoint看作是它的继承类或子类。
//不过这种方式继承的方法并不能实现C++中虚函数的多态特性。
//所有继承来的方法的接收者参数依然是那个匿名成员本身，而不是当前的变量。
type ColoredPoint struct {
	Point
	Color color.RGBA
}



func main() {

}
