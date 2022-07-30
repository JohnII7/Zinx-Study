package znet

import (
	"Zinx_Study/zinx/utils"
	"Zinx_Study/zinx/ziface"
	"bytes"
	"encoding/binary"
	"errors"
)

type DataPack struct {
}

// 拆包封包实例的一个初始化方法
func NewDataPack() *DataPack {
	return &DataPack{}
}

// 获取包的头长度
func (dp *DataPack) GetHeadLen() uint32 {
	// DataLen uint32(4字节) + ID uint32(4字节)
	return 8
}

// 封包方法
func (dp *DataPack) Pack(msg ziface.IMessage) ([]byte, error) {
	// 创建一个存放bytes字节缓冲
	dataBuf := bytes.NewBuffer([]byte{})

	// 将DataLen写进dataBuff中
	if err := binary.Write(dataBuf, binary.LittleEndian, msg.GetMsgLen()); err != nil {
		return nil, err
	}

	// 将MsgId写进dataBuff中
	if err := binary.Write(dataBuf, binary.LittleEndian, msg.GetMsgId()); err != nil {
		return nil, err
	}

	// 将data数据写进dataBuff中
	if err := binary.Write(dataBuf, binary.LittleEndian, msg.GetData()); err != nil {
		return nil, err
	}

	return dataBuf.Bytes(), nil
}

// 拆包方法(将包的head信息读出来) 之后再根据head信息里的data长度，再进行一次读
func (dp *DataPack) Unpack(binaryData []byte) (ziface.IMessage, error) {
	// 创建一个从输入二进制数据的ioReader
	dataBuf := bytes.NewReader(binaryData)

	// 只解压head信息, 得到dataLen和MsgId
	msg := &Message{}

	// 读dataLen
	if err := binary.Read(dataBuf, binary.LittleEndian, &msg.DataLen); err != nil {
		return nil, err
	}

	// 读MsgId
	if err := binary.Read(dataBuf, binary.LittleEndian, &msg.Id); err != nil {
		return nil, err
	}

	// 判断dataLen是否已经超出允许全局允许的最大长度
	if utils.GlobalObject.MaxPackageSize > 0 && msg.DataLen > utils.GlobalObject.MaxPackageSize {
		return nil, errors.New("too large msg data receive")
	}

	return msg, nil
}
