package main

import (
	"net"
	"net/rpc"
	"time"
)

type HelloService struct {}
func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}



func main() {
	rpc.RegisterName("HelloService", new(HelloService))
	//以下是启动反向RPC服务的代码：
	//todo 反向RPC的内网服务将不再主动提供TCP监听服务，而是首先主动链接到对方的TCP服务器。
	// 然后基于每个建立的TCP链接向对方提供RPC服务。
	for {
		conn, _ := net.Dial("tcp", "localhost:1234")
		if conn == nil {
			time.Sleep(time.Second)
			continue
		}

		rpc.ServeConn(conn)
		conn.Close()
	}
	//listener, err := net.Listen("tcp", ":1234")
	//if err != nil {
	//	log.Fatal("ListenTCP error:", err)
	//}
	//
	//conn, err := listener.Accept()
	//if err != nil {
	//	log.Fatal("Accept error:", err)
	//}
	//
	//rpc.ServeConn(conn)
}






