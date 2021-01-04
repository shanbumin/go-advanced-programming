package main

import (
	"io"
	"log"
	"net/http"
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



//todo 注意这里的http服务又承担了rpc客户端以及服务端的职责额
func main() {

	rpc.RegisterName(HelloServiceName, new(HelloService))

	http.HandleFunc("/jsonrpc", func(w http.ResponseWriter, r *http.Request) {
		//基于http.ResponseWriter和http.Request类型的参数构造一个io.ReadWriteCloser类型的conn通道。
		var conn io.ReadWriteCloser = struct {
			io.Writer
			io.ReadCloser
		}{
			ReadCloser: r.Body,
			Writer:     w,
		}
		//然后基于conn构建针对服务端的json编码解码器。
		//最后通过rpc.ServeRequest函数为每次请求处理一次RPC方法调用。
		rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
	})

	err:=http.ListenAndServe(":1234", nil)
	if err !=nil{
		log.Fatal(err)
	}
}


//curl localhost:1234/jsonrpc -X POST --data '{"method":"ch4/rpc/test.HelloService.Hello","params":["hello"],"id":0}'


