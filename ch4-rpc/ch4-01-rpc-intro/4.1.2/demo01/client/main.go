package main

import (
	"fmt"
	"log"
	"net/rpc"
)
const HelloServiceName = "ch4/rpc/test.HelloService"
//首先是通过rpc.Dial拨号RPC服务，然后通过client.Call调用具体的RPC方法。
//在调用client.Call时，第一个参数是用点号链接的RPC服务名字和方法名字，第二和第三个参数分别我们定义RPC方法的两个参数。
func main() {


	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	//其中唯一的变化是client.Call的第一个参数用HelloServiceName+".Hello"代替了"HelloService.Hello"。
	//然而通过client.Call函数调用RPC方法依然比较繁琐，同时参数的类型依然无法得到编译器提供的安全保障。
	//todo 为了简化客户端用户调用RPC函数，我们在可以在接口规范部分增加对客户端的简单包装：
	var reply string
	err = client.Call(HelloServiceName+".Hello", "hello", &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)
}