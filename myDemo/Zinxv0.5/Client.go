package main

import (
	"Zinx_Study/zinx/znet"
	"fmt"
	"io"
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
		// 发送封包message	MsgID: 0
		dp := znet.NewDataPack()
		binaryMsg, err := dp.Pack(znet.NewMsgPackage(0, []byte("Zinxv0.5 client Test Message")))
		if err != nil {
			fmt.Println("Pack error:", err)
			return
		}
		if _, err := conn.Write(binaryMsg); err != nil {
			fmt.Println("write error:", err)
			return
		}
		// 服务器此时应该回复一个message数据, MsgID: 1 ping...ping...ping

		// 先读取流中head部分, 得到ID和 dataLen
		binaryHead := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(conn, binaryHead); err != nil {
			fmt.Println("read head error:", err)
			break
		}

		// 将二进制的head拆包到msg结构体中
		msgHead, err := dp.Unpack(binaryHead)
		if err != nil {
			fmt.Println("client unpackmsgHead error:", err)
			break
		}
		if msgHead.GetMsgLen() > 0 {
			// 再根据dataLen进行第二次读取, 将data读出来
			// msg := znet.Message
			msg := msgHead.(*znet.Message)
			msg.Data = make([]byte, msg.GetMsgLen())

			if _, err := io.ReadFull(conn, msg.Data); err != nil {
				fmt.Println("read msg data error:", err)
				return
			}

			fmt.Println("--->Receive Server Msg: ID =", msg.Id, "len =", msg.DataLen, ",data =", string(msg.Data))
		}

		// cpu阻塞
		time.Sleep(time.Second * 1)

	}

	// 连接调用Write写数据
}
