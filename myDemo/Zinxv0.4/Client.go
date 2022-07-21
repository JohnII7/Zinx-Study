package main

import (
	"fmt"
	"net"
	"time"
)

// 模拟客户端
func main() {
	fmt.Println(">>>client start<<<")
	time.Sleep(time.Second * 1)

	// 直接连接远程服务器, 得到一个conn连接
	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("client start err, exit")
		return
	}
	for {
		_, err := conn.Write([]byte("Hello Zinx V.02"))
		if err != nil {
			fmt.Println("writer connection err:", err)
		}
		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read buf error")
			return
		}
		fmt.Printf("server call back: %s, cnt: %d\n", buf, cnt)

		// cpu阻塞
		time.Sleep(time.Second * 1)

	}

	// 连接调用Write写数据
}
