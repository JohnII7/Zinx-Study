package znet

import (
	"Zinx_Study/zinx/ziface"
	"errors"
	"fmt"
	"net"
)

// Server IServer接口的实现, 定义一个Server的服务器模块
type Server struct {
	// 服务器名称
	Name string

	// 服务器绑定的ip版本
	IPVersion string

	// 服务器监听的IP
	IP string

	// 端口
	Port int
}

// CallBackToClient 定义当前客户端连接的所绑定的 handle api(目前handle是写死的, 以后优化应该由用户自定义handle方法)
func CallBackToClient(conn *net.TCPConn, data []byte, cnt int) error {
	// 回显的业务
	fmt.Println("[Conn Handle] CallBackToClient...")
	if _, err := conn.Write(data[:cnt]); err != nil {
		fmt.Println("write back buf err", err)
		return errors.New("CallBackToClient")
	}
	return nil
}

// Start 服务器
func (s *Server) Start() {
	fmt.Printf("[Start] Server Listener at IP:%s, Port: %d is starting\n", s.IP, s.Port)
	go func() {
		// 获取一个TCP的addr
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr error", err)
			return
		}

		// 监听服务器的地址
		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen ", s.IPVersion, " err", err)
			return
		}
		fmt.Println(">>>start Zinx server success, ", s.Name, " Listening<<<")
		var cid uint32 = 0
		// 阻塞等待客户端连接, 处理客户端连接(读写)
		for {
			// 如果有客户端连接过来, 阻塞返回
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err", err)
				continue
			}
			// 将处理新连接的业务方法和 conn 进行绑定, 得到连接模块
			delConn := NewConnection(conn, cid, CallBackToClient)
			cid++

			// 启动当前连接业务
			go delConn.Start()
		}
	}()
}

// Stop 服务器
func (s *Server) Stop() {
	// TODO 将一些服务器的资源、状态或者一些已经开辟的连接信息进行停止或者回收
}

// Server 运行服务器
func (s *Server) Server() {
	// 启动Server服务
	s.Start()

	// TODO 做一些启动服务器之后的额外业务

	// 阻塞
	select {}
}

// NewServer 初始化Server模块
func NewServer(name string) ziface.ISever {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8999,
	}
	return s
}
