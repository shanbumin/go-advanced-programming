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

	//Go语言的RPC库最简单的使用方式是通过`Client.Call`方法进行同步阻塞调用，该方法的实现如下：
	//todo 我们也可以通过`Client.Go`方法异步调用HelloService服务：
	//err = client.Call("HelloService.Hello", "hello", &reply)
	//Go方法中`client.send`方法调用是线程安全的，因此可以从多个Goroutine同时向同一个RPC链接发送调用指令。
	helloCall := client.Go("HelloService.Hello", "hello", new(string), nil)
	helloCall = <-helloCall.Done
	if err := helloCall.Error; err != nil {
		log.Fatal(err)
	}

	// do some thing
	//在异步调用命令发出后，一般会执行其他的任务，因此异步调用的输入参数和返回值可以通过返回的Call变量进行获取。
	args := helloCall.Args.(string)
	reply := helloCall.Reply.(*string)
	fmt.Println(args, *reply)

}


/*



func (client *Client) Call(serviceMethod string, args interface{}, reply interface{}) error {
	call := <-client.Go(serviceMethod, args, reply, make(chan *Call, 1)).Done
	return call.Error
}





 */