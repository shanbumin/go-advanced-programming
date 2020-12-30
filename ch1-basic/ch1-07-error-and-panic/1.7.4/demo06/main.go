package main

import "fmt"
import "errors"

//当希望将捕获到的异常转为错误时，如果希望忠实返回原始的信息，需要针对不同的类型分别处理：
func main() {
	err:=foo()
	if err !=nil{
		fmt.Println(err)
	}
}



func foo() (err error) {
	defer func() {
		if r := recover(); r != nil {
			switch x := r.(type) {
			case string:
				err = errors.New(x)
			case error:
				err = x
			default:
				err = fmt.Errorf("Unknown panic: %v", r)
			}
		}
	}()

	panic("TODO")
}


//todo 不过这样做和Go语言简单直接的编程哲学背道而驰了。

