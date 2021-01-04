package main

import (
	"fmt"
	"log"
	"net/rpc"
)



//todo String为输入输出参数
func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var reply *String
	request:=&String{
		Value:"sam",
	}
	err = client.Call("HelloService.Hello",request, &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)
}