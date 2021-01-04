package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)
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

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}
		//go rpc.ServeConn(conn)
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn)) //针对服务器端的JSON编解码器
		//todo 代码中最大的变化是用rpc.ServeCodec函数替代了rpc.ServeConn函数，传入的参数是针对服务端的json编解码器。
	}

}