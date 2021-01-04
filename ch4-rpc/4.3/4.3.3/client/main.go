package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

//首先从管道去取一个RPC客户端对象，并且通过defer语句指定在函数退出前关闭客户端。然后是执行正常的RPC调用。
func doClientWork(clientChan <-chan *rpc.Client) {
	client := <-clientChan
	defer client.Close()

	var reply string
	err := client.Call("HelloService.Hello", "hello", &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)
}

//而RPC客户端则需要在一个公共的地址提供一个TCP服务，用于接受RPC服务器的链接请求：
//当每个链接建立后，基于网络链接构造RPC客户端对象并发送到clientChan管道。
//客户端执行RPC调用的操作在doClientWork函数完成：

func main() {

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	clientChan := make(chan *rpc.Client)

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Fatal("Accept error:", err)
			}

			clientChan <- rpc.NewClient(conn)
		}
	}()

	doClientWork(clientChan)



	//client, err := rpc.Dial("tcp", "localhost:1234")
	//if err != nil {
	//	log.Fatal("dialing:", err)
	//}
	//
	//var reply string
	////Go语言的RPC库最简单的使用方式是通过`Client.Call`方法进行同步阻塞调用，该方法的实现如下：
	////todo 可以跳转进去查看Call的具体实现
	////首先通过`Client.Go`方法进行一次异步调用，返回一个表示这次调用的`Call`结构体。然后等待`Call`结构体的Done管道返回调用结果。
	////我们也可以通过`Client.Go`方法异步调用HelloService服务：
	//err = client.Call("HelloService.Hello", "hello", &reply)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println(reply)
}

