package main

import (
	"Zinx_Study/zinx/ziface"
	"Zinx_Study/zinx/znet"
	"fmt"
)

/**
基于Zinx框架来开发的服务端应用程序
*/
// ping测试 自定义路由
type PingRouter struct {
	znet.BaseRouter
}

// Test PreHandle
func (this *PingRouter) PreHandle(request ziface.IRequest) {
	fmt.Println("Call Router PreHandle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte(">>>before ping<<<\n"))
	if err != nil {
		fmt.Println("call back before ping error")
	}

}

// Test Handle
func (this *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call Router Handle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte(">>>ping ping<<<\n"))
	if err != nil {
		fmt.Println("call back ping error")
	}
}

// Test PostHandle
func (this *PingRouter) PostHandle(request ziface.IRequest) {
	fmt.Println("Call Router PostHandle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte(">>>after ping<<<\n"))
	if err != nil {
		fmt.Println("call back after ping error")
	}
}

func main() {
	// 创建一个Server句柄, 使用Zinx的api
	s := znet.NewServer("[zinx V0.4]")

	//给当前Zinx框架添加一个自定义的Router
	s.AddRouter(&PingRouter{})

	//启动Server
	s.Server()
}
