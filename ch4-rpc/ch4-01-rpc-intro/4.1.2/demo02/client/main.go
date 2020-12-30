package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type HelloServiceClient struct {
	*rpc.Client
}

//var _ HelloServiceInterface = (*HelloServiceClient)(nil)

func DialHelloService(network, address string) (*HelloServiceClient, error) {
	c, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &HelloServiceClient{Client: c}, nil
}

func (p *HelloServiceClient) Hello(request string, reply *string) error {
	return p.Client.Call(HelloServiceName+".Hello", request, reply)
}

//---------------------------------------------------------------------------------
const HelloServiceName = "ch4/rpc/test.HelloService"

func main() {
	client, err := DialHelloService("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var reply string
	//现在客户端用户不用再担心RPC方法名字或参数类型不匹配等低级错误的发生。
	err = client.Hello("hello", &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
}