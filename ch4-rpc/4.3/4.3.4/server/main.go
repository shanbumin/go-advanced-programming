package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

//首先改造HelloService，里面增加了对应链接的conn成员：
//基于上下文信息，我们可以方便地为RPC服务增加简单的登陆状态的验证：
type HelloService struct {
	conn net.Conn
	isLogin bool
}
//Hello方法中就可以根据conn成员识别不同链接的RPC调用：
//func (p *HelloService) Hello(request string, reply *string) error {
//	*reply = "hello:" + request + ", from" + p.conn.RemoteAddr().String()
//	return nil
//}


//func (p *HelloService) Hello(request string, reply *string) error {
//	*reply = "hello:" + request
//	return nil
//}

//这样可以要求在客户端链接RPC服务时，首先要执行登陆操作，登陆成功后才能正常执行其他的服务。
func (p *HelloService) Login(request string, reply *string) error {
	if request != "user:password" {
		return fmt.Errorf("auth failed")
	}
	log.Println("login ok")
	p.isLogin = true
	return nil
}

func (p *HelloService) Hello(request string, reply *string) error {
	if !p.isLogin {
		return fmt.Errorf("please login")
	}
	*reply = "hello:" + request + ", from" + p.conn.RemoteAddr().String()
	return nil
}




//todo 然后为每个链接启动独立的RPC服务：
func main() {

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}

		go func() {
			defer conn.Close()

			p := rpc.NewServer()
			p.Register(&HelloService{conn: conn})
			p.ServeConn(conn)
		} ()
	}


	//rpc.RegisterName("HelloService", new(HelloService))
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






