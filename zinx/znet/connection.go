package znet

import "C"
import (
	"Zinx_Study/zinx/ziface"
	"fmt"
	"net"
)

type Connection struct {
	// 当前连接的socket TCP套接字
	Conn *net.TCPConn

	// 连接的ID
	ConnID uint32

	// 当前连接状态
	isClosed bool

	// 当前连接所绑定的处理业务方法API
	handleAPI ziface.HandleFunc

	// 告知当前连接已经 退出的/停止的 channel
	ExitChan chan bool
}

// NewConnection 初始化连接模块的方法
func NewConnection(conn *net.TCPConn, connID uint32, callbackApi ziface.HandleFunc) *Connection {
	c := &Connection{
		Conn:      conn,
		ConnID:    connID,
		handleAPI: callbackApi,
		isClosed:  false,
		ExitChan:  make(chan bool, 1),
	}
	return c
}

// StartReader 连接读数据业务
func (c *Connection) StartReader() {
	fmt.Println(">>>Reader Goroutine is running<<<")
	defer fmt.Println("connId:", c.ConnID, " Reader is exit, remote addr is", c.RemoteAddr().String())
	defer c.Stop()
	for {
		// 读取客户端的数据到buf中, 最大512字节
		buf := make([]byte, 512)
		cnt, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("receive buf err", err)
			continue
		}
		// 调用当前连接所绑定的HandlerAPI
		if err := c.handleAPI(c.Conn, buf, cnt); err != nil {
			fmt.Println("ConnID:", c, c.ConnID, " handle is error", err)
			break
		}
	}

}

// Start 启动连接 让当前连接准备开始工作
func (c *Connection) Start() {
	fmt.Println(">>>Connection Start()...ConnID:", c.ConnID, "<<<")

	// 启动从当前连接的读数据业务
	go c.StartReader()

	// TODO 启动从当前连接的写数据业务
}

// Stop 停止连接 结束当前连接的工作
func (c *Connection) Stop() {
	fmt.Println(">>>Connection Stop()...ConnID:", c.ConnID, "<<<")

	// 如果当前连接已经关闭
	if c.isClosed == true {
		return
	}
	c.isClosed = true

	// 关闭socket连接
	c.Conn.Close()

	// 回收资源
	close(c.ExitChan)
}

// GetTCPConnection 获取当前连接的绑定socket conn
func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

// GetConnID  获取当前连接模块的连接ID
func (c *Connection) GetConnID() uint32 {
	return c.GetConnID()
}

// RemoteAddr 获取远程客户端的TCP状态 IP port
func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

// Send 发送数据，将数据发送给远程的客户端
func (c *Connection) Send(data []byte) error {
	return nil
}