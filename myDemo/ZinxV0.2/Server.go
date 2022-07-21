package main

import (
	"Zinx_Study/zinx/znet"
)

/**
基于Zinx框架来开发的服务端应用程序
*/

func main() {
	// 创建一个Server句柄, 使用Zinx的api
	s := znet.NewServer("[zinx V0.3]")

	//启动Server
	s.Server()
}
