package main

//切片类型强制转换
func main() {






//
//		为了安全，当两个切片类型`[]T`和`[]Y`的底层原始切片类型不同时，Go语言是无法直接转换类型的。不过安全都是有一定代价的，有时候这种转换是有它的价值的——可以简化编码或者是提升代码的性能。比如在64位系统上，需要对一个`[]float64`切片进行高速排序，我们可以将它强制转为`[]int`整数切片，然后以整数的方式进行排序（因为`float64`遵循IEEE754浮点数标准特性，当浮点数有序时对应的整数也必然是有序的）。
//
//	下面的代码通过两种方法将`[]float64`类型的切片转换为`[]int`类型的切片：
//
//	```go
//// +build amd64 arm64
//
//import "sort"
//
//var a = []float64{4, 2, 5, 7, 2, 1, 88, 1}
//
//func SortFloat64FastV1(a []float64) {
//   // 强制类型转换
//   var b []int = ((*[1 << 20]int)(unsafe.Pointer(&a[0])))[:len(a):cap(a)]
//
//   // 以int方式给float64排序
//   sort.Ints(b)
//}
//
//func SortFloat64FastV2(a []float64) {
//   // 通过 reflect.SliceHeader 更新切片头部信息实现转换
//   var c []int
//   aHdr := (*reflect.SliceHeader)(unsafe.Pointer(&a))
//   cHdr := (*reflect.SliceHeader)(unsafe.Pointer(&c))
//   *cHdr = *aHdr
//
//   // 以int方式给float64排序
//   sort.Ints(c)
//}
//```
//
//	第一种强制转换是先将切片数据的开始地址转换为一个较大的数组的指针，然后对数组指针对应的数组重新做切片操作。中间需要`unsafe.Pointer`来连接两个不同类型的指针传递。需要注意的是，Go语言实现中非0大小数组的长度不得超过2GB，因此需要针对数组元素的类型大小计算数组的最大长度范围（`[]uint8`最大2GB，`[]uint16`最大1GB，以此类推，但是`[]struct{}`数组的长度可以超过2GB）。
//
//	第二种转换操作是分别取到两个不同类型的切片头信息指针，任何类型的切片头部信息底层都是对应`reflect.SliceHeader`结构，然后通过更新结构体方式来更新切片信息，从而实现`a`对应的`[]float64`切片到`c`对应的`[]int`类型切片的转换。
//
//	通过基准测试，我们可以发现用`sort.Ints`对转换后的`[]int`排序的性能要比用`sort.Float64s`排序的性能好一点。不过需要注意的是，这个方法可行的前提是要保证`[]float64`中没有NaN和Inf等非规范的浮点数（因为浮点数中NaN不可排序，正0和负0相等，但是整数中没有这类情形）。



}
