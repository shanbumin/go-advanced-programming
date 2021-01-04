package main

import (
	"fmt"
	"log"
	"net/rpc"
)



func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var reply string
	//Go语言的RPC库最简单的使用方式是通过`Client.Call`方法进行同步阻塞调用，该方法的实现如下：
	//todo 可以跳转进去查看Call的具体实现
	//首先通过`Client.Go`方法进行一次异步调用，返回一个表示这次调用的`Call`结构体。然后等待`Call`结构体的Done管道返回调用结果。
	//我们也可以通过`Client.Go`方法异步调用HelloService服务：
	err = client.Call("HelloService.Hello", "hello", &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)
}


/*



func (client *Client) Call(serviceMethod string, args interface{}, reply interface{}) error {
	call := <-client.Go(serviceMethod, args, reply, make(chan *Call, 1)).Done
	return call.Error
}





 */