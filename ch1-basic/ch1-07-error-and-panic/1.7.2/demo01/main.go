package main

import (
	"github.com/chai2010/errors"
	"io/ioutil"
)

func loadConfig() error {
	_, err := ioutil.ReadFile("/path/to/file")
	if err != nil {
		return errors.Wrap(err, "read failed")
	}
	return  nil
}


func setup() error {
	err := loadConfig()
	if err != nil{
		return errors.Wrap(err, "invalid config")
	}
	return  nil
}



func main() {
		if err := setup(); err != nil {
		    //log.Fatal(err)
			//1.上面的例子中，错误被进行了2层包装。我们可以这样遍历原始错误经历了哪些包装流程：

			//for i, e := range err.(errors.Error).Wraped() {
			//	fmt.Printf("wraped(%d): %v\n", i, e)
			//}


			//2.同时也可以获取每个包装错误的函数调用堆栈信息：
			//for i, x := range err.(errors.Error).Caller() {
			//	fmt.Printf("caller:%d: %s\n", i, x.FuncName)
			//}





		}




}