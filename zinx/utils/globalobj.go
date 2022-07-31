package utils

import (
	"Zinx_Study/zinx/ziface"
	"encoding/json"
	"io/ioutil"
)

/**
存储一切有关Zinx框架的全局参数, 供其他模块使用
一些参数可以通过zinx.json由用户全局配置
*/

type GlobalObj struct {
	// server
	TCPServer ziface.ISever // 当前Zinx全局的Server对象
	Host      string        // 当前服务器主机监听的IP
	TcpPort   int           // 当前服务器主机监听的端口号
	Name      string        // 当前服务器名称

	// Zinx
	Version        string // 当前Zinx的版本号
	MaxConn        int    //当前服务器允许的最大连接数
	MaxPackageSize uint32 // 当前Zinx框架数据包的最大值

}

var GlobalObject *GlobalObj

// 从Zinx.json去加载用于自定义的参数
func (g *GlobalObj) Reload() {
	data, err := ioutil.ReadFile("conf/zinx.json")

	if err != nil {
		panic(err)
	}
	// 将json文件解析到strut中
	err = json.Unmarshal(data, &GlobalObject)
	if err != nil {
		panic(err)
	}
}

// 提供init方法, 初始化当前的GlobalObject
func init() {
	// 如果配置未见没有加载, 默认的值
	GlobalObject = &GlobalObj{
		Name:           "ZinxServerApp",
		Version:        "v0.5",
		TcpPort:        8999,
		Host:           "0.0.0.0",
		MaxConn:        1000,
		MaxPackageSize: 4096,
	}

	// 应该尝试从conf/zinx.json去加载一些用户自定义的参数
	GlobalObject.Reload()
}
