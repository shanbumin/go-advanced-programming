package main

import (
	"log"
	"net"
	"net/rpc"
	"shanbumin/go-advanced-programming/ch4-rpc/4.3/4.3.2/kv"
)


func main() {
	rpc.RegisterName("KVStoreService", new(kv.KVStoreService))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("Accept error:", err)
	}

	rpc.ServeConn(conn)
}
