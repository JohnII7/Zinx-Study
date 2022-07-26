package znet

import "C"
import (
	"Zinx_Study/zinx/ziface"
	"errors"
	"fmt"
	"io"
	"net"
)

type Connection struct {
	// 当前连接的socket TCP套接字
	Conn *net.TCPConn

	// 连接的ID
	ConnID uint32

	// 当前连接状态
	isClosed bool

	// 告知当前连接已经 退出的/停止的 channel
	ExitChan chan bool

	// 该链接处理的方法Router
	Router ziface.IRouter
}

// NewConnection 初始化连接模块的方法
func NewConnection(conn *net.TCPConn, connID uint32, router ziface.IRouter) *Connection {
	c := &Connection{
		Conn:     conn,
		ConnID:   connID,
		Router:   router,
		isClosed: false,
		ExitChan: make(chan bool, 1),
	}
	return c
}

// StartReader 连接读数据业务
func (c *Connection) StartReader() {
	fmt.Println(">>>Reader Goroutine is running<<<")
	defer fmt.Println("connId:", c.ConnID, " Reader is exit, remote addr is", c.RemoteAddr().String())
	defer c.Stop()
	for {
		// 读取客户端的数据到buf中, 最大MaxPackageSize
		/*buf := make([]byte, utils.GlobalObject.MaxPackageSize)
		_, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("receive buf err", err)
			continue
		}*/

		// 创建一个拆包解包对象
		dp := NewDataPack()

		// 读取客户端的Msg Head 二进制流 8个字节,
		headData := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(c.GetTCPConnection(), headData); err != nil {
			fmt.Println("read msg head error:", err)
			break
		}

		// 拆包, 得到msgID 和 msgDataLen 放在msg消息中
		msg, err := dp.Unpack(headData)
		if err != nil {
			fmt.Println("unpack error:", err)
			break
		}

		// 根据dataLen 再次读取Data, 放在msg.Data中
		var data []byte
		if msg.GetMsgLen() > 0 {
			data = make([]byte, msg.GetMsgLen())
			if _, err := io.ReadFull(c.GetTCPConnection(), data); err != nil {
				fmt.Println("read msg data error", err)
				break
			}
		}
		msg.SetData(data)

		// 得到当前conn数据的Request请求数据
		req := Request{
			conn: c,
			msg:  msg,
		}
		go func(request ziface.IRequest) {
			c.Router.PreHandle(request)
			c.Router.Handle(request)
			c.Router.PostHandle(request)
		}(&req)
		// 从路由中找到注册绑定的Conn对用router调用
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

// 提供一个SendMsg方法 将我们要发送给客户端的数据先进行封包,再发送
func (c *Connection) SendMsg(msgId uint32, data []byte) error {
	if c.isClosed == true {
		return errors.New("Connection closed when send msg")
	}

	// 将数据进行封包 MsgDataLen|MsgID|Data
	dp := NewDataPack()

	binaryMsg, err := dp.Pack(NewMsgPackage(msgId, data))
	if err != nil {
		fmt.Println("Pack error msg id:", msgId)
		return errors.New("Pack error msg")
	}
	// 将数据发送给客户端
	if _, err := c.Conn.Write(binaryMsg); err != nil {
		fmt.Println("write msg id:", msgId, "error:", err)
		return errors.New("conn write error")
	}

	return nil
}
