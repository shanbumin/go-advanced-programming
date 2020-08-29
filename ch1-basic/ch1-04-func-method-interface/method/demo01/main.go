package main



//1.Go语言的祖先C语言虽然不是一个支持面向对象的语言，但是C语言的标准库中的File相关的函数也用到了的面向对象编程的思想。下面我们实现一组C语言风格的File函数：
//其中`OpenFile`类似构造函数用于打开文件对象，
//`CloseFile`类似析构函数用于关闭文件对象，`ReadFile`则类似普通的成员函数，这三个函数都是普通的函数。
//`CloseFile`和`ReadFile`作为普通函数，需要占用包级空间中的名字资源。
// 不过`CloseFile`和`ReadFile`函数只是针对`File`类型对象的操作，这时候我们更希望这类函数和操作对象的类型紧密绑定在一起。


// 文件对象
type File struct {
	fd int
}

// 打开文件
func OpenFile(name string) (f *File, err error) {
	return
}

// 关闭文件
func CloseFile(f *File) error {
	return nil
}

// 读文件数据
func ReadFile(f *File, offset int64, data []byte) int {
	return 0
}

//2.Go语言中的做法是，将`CloseFile`和`ReadFile`函数的第一个参数移动到函数名的开头：
// 关闭文件
func (f *File) CloseFile() error {
	// ...
	return nil
}

// 读文件数据
func (f *File) ReadFile(offset int64, data []byte) int {
	// ...
	return  0
}
//这样的话，`CloseFile`和`ReadFile`函数就成了`File`类型独有的方法了（而不是`File`对象方法）。
//它们也不再占用包级空间中的名字资源，同时`File`类型已经明确了它们操作对象，因此方法名字一般简化为`Close`和`Read`：

// 关闭文件
func (f *File) Close() error {
   // ...
	return  nil
}

// 读文件数据
func (f *File) Read(offset int64, data []byte) int {
   // ...
	return 0
}




func main() {


	
}
