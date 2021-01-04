package main

import (
	"log"
	"net"
	"net/rpc"
)


//RPC服务的接口规范分为三个部分：
//首先是服务的名字，
//然后是服务要实现的详细方法列表，
//最后是注册该类型服务的函数。
//为了避免名字冲突，我们在RPC服务的名字中增加了包路径前缀（这个是RPC服务抽象的包路径，并非完全等价Go语言的包路径）。
//RegisterHelloService注册服务时，编译器会要求传入的对象满足HelloServiceInterface接口。

const HelloServiceName = "ch4/rpc/test.HelloService"

type HelloServiceInterface = interface {
	Hello(request string, reply *string) error
}

func RegisterHelloService(svc HelloServiceInterface) error {
	return rpc.RegisterName(HelloServiceName, svc)
}
//------------------------------------
type HelloService struct {}
func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}


func main() {
	RegisterHelloService(new(HelloService))
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("Accept error:", err)
	}

	rpc.ServeConn(conn)
}
