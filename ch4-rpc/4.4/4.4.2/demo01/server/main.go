package main

import (
	"context"
	"log"
	"net"
	"google.golang.org/grpc"

)

//基于服务端的HelloServiceServer接口可以重新实现HelloService服务：
type HelloServiceImpl struct{}

func (p *HelloServiceImpl) Hello(ctx context.Context, args *String) (*String, error) {
	reply := &String{Value: "hello:" + args.GetValue()}
	return reply, nil
}


//gRPC服务的启动流程和标准库的RPC服务启动流程类似：

//首先是通过`grpc.NewServer()`构造一个gRPC服务对象，
//然后通过gRPC插件生成的RegisterHelloServiceServer函数注册我们实现的HelloServiceImpl服务。
//然后通过`grpcServer.Serve(lis)`在一个监听端口上提供gRPC服务。

func main() {
	grpcServer := grpc.NewServer()
	RegisterHelloServiceServer(grpcServer, new(HelloServiceImpl))

	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer.Serve(lis)
}
