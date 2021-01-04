package main

import (
	"log"
	"net"
	"net/rpc"
)


type HelloService struct {}
//func (p *HelloService) Hello(request string, reply *string) error {
func (p *HelloService) Hello(request *String, reply *String) error {
	//*reply = "hello:" + request
	reply.Value ="hello:"+request.GetValue()
	return nil
}

//todo 这里我们让RPC与Protobuf简单的结合使用了
func main() {
	//将HelloService类型的对象注册为一个RPC服务
	rpc.RegisterName("HelloService", new(HelloService))

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
