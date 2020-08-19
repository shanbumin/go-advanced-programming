package main

import (
	"fmt"
	"os"
	"github.com/bwmarrin/snowflake"
)

func main() {
	//使用唯一的节点号构造一个新的雪花节点。默认设置允许节点号范围为0到1023。如果您设置了自定义NodeBits值，则需要计算节点号范围。
	//创建一个节点号为1的新节点
	n, err := snowflake.NewNode(1)
	if err != nil {
		println(err)
		os.Exit(1)
	}

	//for i := 0; i < 3; i++ {
		//使用节点对象，调用Generate()方法来生成并返回唯一的雪花ID。
		id := n.Generate()


		//以几种不同的方式打印出ID。
		fmt.Printf("Int64  ID: %d\n", id)
		fmt.Printf("String ID: %s\n", id)
		fmt.Printf("Base2  ID: %s\n", id.Base2())
		fmt.Printf("Base64 ID: %s\n", id.Base64())
		fmt.Println("---------------------------------------")


		//打印ID的时间戳
		fmt.Printf("ID Time  : %d\n", id.Time())
		//打印ID的节点号
		fmt.Printf("ID Node  : %d\n", id.Node())
		//打印ID的序列号
		fmt.Printf("ID Step  : %d\n", id.Step())
		fmt.Println("---------------------------------------")
		//一次生成和打印。
		fmt.Printf("ID       : %d\n", n.Generate().Int64())



	//}
}