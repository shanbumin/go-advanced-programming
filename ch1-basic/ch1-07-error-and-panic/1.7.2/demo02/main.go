package main

import (
	"fmt"
	"github.com/chai2010/errors"
)

func main() {


	//对于基于http协议的网络服务，我们还可以给错误绑定一个对应的http状态码：
	err := errors.NewWithCode(404, "http error code")

	fmt.Println(err)
	fmt.Println(err.(errors.Error).Code())


}
