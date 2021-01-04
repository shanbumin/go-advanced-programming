package main

import (
	"fmt"
	"log"
	"net/rpc"
	"time"
)


//首先启动一个独立的Goroutine监控key的变化。
//同步的watch调用会阻塞，直到有key发生变化或者超时。
//然后在通过Set方法修改KV值时，服务器会将变化的key通过Watch方法返回。
//这样我们就可以实现对某些状态的监控。

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	go func() {
		var keyChanged string
		err := client.Call("KVStoreService.Watch", 30, &keyChanged)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("watch:", keyChanged)
	} ()

	err = client.Call("KVStoreService.Set", [2]string{"abc", "abc-value"}, new(struct{}),
	)
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(time.Second*3)


	//var reply string
	//err = client.Call("HelloService.Hello", "hello", &reply)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println(reply)
}

