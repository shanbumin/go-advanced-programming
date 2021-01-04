package main

import (
	"log"
	"net"
	"net/rpc"
	message "shanbumin/go-advanced-programming/ch4-rpc/4.2/4.2.1/demo02"
	"time"
	"errors"
)



//订单服务
type OrderService struct {
}
func (os *OrderService) GetOrderInfo(request message.OrderRequest, response *message.OrderInfo) error {
	//201907310003
	orderMap := map[string]message.OrderInfo{
		"201907300001": message.OrderInfo{OrderId: "201907300001", OrderName: "衣服", OrderStatus: "已付款"},
		"201907310001": message.OrderInfo{OrderId: "201907310001", OrderName: "零食", OrderStatus: "已付款"},
		"201907310002": message.OrderInfo{OrderId: "201907310002", OrderName: "食品", OrderStatus: "未付款"},
	}
	current := time.Now().Unix()
	if (request.TimeStamp > current) {
		*response = message.OrderInfo{OrderId: "0", OrderName: "", OrderStatus: "订单信息异常"}
	} else {
		result := orderMap[request.OrderId]//201907310003
		if result.OrderId != "" {
			*response = orderMap[request.OrderId]
		} else {
			return errors.New("server error")
		}
	}
	return nil
}



func main() {
	//将OrderService类型的对象注册为一个RPC服务
	rpc.RegisterName("OrderService", new(OrderService))

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
