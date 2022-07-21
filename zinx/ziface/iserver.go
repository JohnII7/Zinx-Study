package ziface

// ISever 定义一个服务器接口
type ISever interface {
	// Start 启动服务器
	Start()

	// Stop 停止服务器
	Stop()

	// Server 运行服务器
	Server()

	// 路由功能: 给当前的服务注册一个路由, 供客户端的连接处理使用
	AddRouter(router IRouter)
}
