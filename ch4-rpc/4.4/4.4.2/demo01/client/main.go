package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

//其中grpc.Dial负责和gRPC服务建立链接，
//然后NewHelloServiceClient函数基于已经建立的链接构造HelloServiceClient对象。
//返回的client其实是一个HelloServiceClient接口对象，通过接口定义的方法就可以调用服务端对应的gRPC服务提供的方法。
//todo gRPC和标准库的RPC框架有一个区别，gRPC生成的接口并不支持异步调用。
//todo 不过我们可以在多个Goroutine之间安全地共享gRPC底层的HTTP/2链接，因此可以通过在另一个Goroutine阻塞调用的方式模拟异步调用。
func main() {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := NewHelloServiceClient(conn)
	reply, err := client.Hello(context.Background(), &String{Value: "hello"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply.GetValue())
}


