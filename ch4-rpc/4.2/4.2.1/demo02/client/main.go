package main

import (
	"fmt"
	"log"
	"net/rpc"
	message "shanbumin/go-advanced-programming/ch4-rpc/4.2/4.2.1/demo02"
	"time"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	timeStamp := time.Now().Unix()
	request := message.OrderRequest{OrderId: "201907310001", TimeStamp: timeStamp}
	var response *message.OrderInfo

	err = client.Call("OrderService.GetOrderInfo",request, &response)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(*response)
}