package main

import "fmt"

func main() {

	var a [3]int                    // 定义长度为3的int型数组, 元素全部为0
	var b = [...]int{1, 2, 3}       // 定义长度为3的int型数组, 元素为 1, 2, 3
	var c = [...]int{2: 3, 1: 2}    // 定义长度为3的int型数组, 元素为 0, 2, 3 (指明下标为2的元素是3，下标为1的元素的是2)

	//我们可以用`for`循环来迭代数组。下面常见的几种方式都可以用来遍历数组：
	//用`for range`方式迭代的性能可能会更好一些，因为这种迭代可以保证不会出现数组越界的情形，每轮迭代对数组元素的访问时可以省去对下标越界的判断。
	for i := range a {
		fmt.Printf("a[%d]: %d\n", i, a[i])
	}
	for i, v := range b {
		fmt.Printf("b[%d]: %d\n", i, v)
	}
	for i := 0; i < len(c); i++ {
		fmt.Printf("c[%d]: %d\n", i, c[i])
	}


	//用`for range`方式迭代，还可以忽略迭代时的下标:
	//其中`times`对应一个`[5][0]int`类型的数组，虽然第一维数组有长度，但是数组的元素`[0]int`大小是0，因此整个数组占用的内存大小依然是0。
	//没有付出额外的内存代价，我们就通过`for range`方式实现了`times`次快速迭代。
   var times [5][0]int
   for range times {
      fmt.Println("hello")
   }



}
