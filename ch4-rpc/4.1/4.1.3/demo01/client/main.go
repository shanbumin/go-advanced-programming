package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

const HelloServiceName = "ch4/rpc/test.HelloService"
//先手工调用net.Dial函数建立TCP链接
//然后基于该链接建立针对客户端的json编解码器。
func main() {
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("net.Dial:", err)
	}
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	var reply string
	err = client.Call(HelloServiceName+".Hello", "hello", &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
}


//todo 如果想看下客户端发什么内容给服务端了，则我们可以通过nc  -l  1234模拟tcp服务端
//  然后再次执行一次RPC客户端调用将会发现nc输出了以下的信息：
//{"method":"ch4/rpc/test.HelloService.Hello","params":["hello"],"id":0}
//这是一个json编码的数据，其中method部分对应要调用的rpc服务和方法组合成的名字，params部分的第一个元素为参数，id是由调用端维护的一个唯一的调用编号。


//请求的json数据对象在内部对应两个结构体：客户端是clientRequest，服务端是serverRequest。
//clientRequest和serverRequest结构体的内容基本是一致的：
/*
type clientRequest struct {
   Method string         `json:"method"`
   Params [1]interface{} `json:"params"`
   Id     uint64         `json:"id"`
}
type serverRequest struct {
   Method string           `json:"method"`
   Params *json.RawMessage `json:"params"`
   Id     *json.RawMessage `json:"id"`
}
------------------------------------------------------------------------------------------------------------------------
*/



//todo 既然我们知道了tcp客户端往服务端发送了什么，我们完全可以模拟客户端直接发送的
//  echo -e '{"method":"ch4/rpc/test.HelloService.Hello","params":["hello"],"id":1}' | nc localhost 1234
//返回的结果也是一个json格式的数据：
//{"id":1,"result":"hello:hello","error":null}
//其中id对应输入的id参数，result为返回的结果，error部分在出问题时表示错误信息。
//对于顺序调用来说，id不是必须的。但是Go语言的RPC框架支持异步调用，当返回结果的顺序和调用的顺序不一致时，可以通过id来识别对应的调用。

//返回的json数据也是对应内部的两个结构体：客户端是clientResponse，服务端是serverResponse。两个结构体的内容同样也是类似的：
/*
type clientResponse struct {
   Id     uint64           `json:"id"`
   Result *json.RawMessage `json:"result"`
   Error  interface{}      `json:"error"`
}

type serverResponse struct {
   Id     *json.RawMessage `json:"id"`
   Result interface{}      `json:"result"`
   Error  interface{}      `json:"error"`
}
*/


//todo  因此无论采用何种语言，只要遵循同样的json结构，以同样的流程就可以和Go语言编写的RPC服务进行通信。这样我们就实现了跨语言的RPC。





