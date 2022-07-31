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

// Test Handle
func (this *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call Router Handle")
	// 先读取客户端的数据, 再回写ping..ping业务

	fmt.Println("receive from client: msgID:", request.GetMsgID(), ", data:", string(request.GetData()))

	err := request.GetConnection().SendMsg(1, []byte("ping...ping...ping"))
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	// 创建一个Server句柄, 使用Zinx的api
	s := znet.NewServer("[zinx V0.5]")

	//给当前Zinx框架添加一个自定义的Router
	s.AddRouter(&PingRouter{})

	//启动Server
	s.Server()
}
