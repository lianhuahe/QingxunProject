package utils

import (
	serveInterface "Sample_Douyin_Server/server/serveinterface"
	"encoding/json"
	"io/ioutil"
)

type GlobalObj struct {
	TcpServer      serveInterface.IServer
	Host           string
	TcpPort        int
	MsgChanLen     int //每个链接读写管道的缓存大小
	MaxConn        int //最大连接数
	WorkerPoolSize int //工作池大小
	WorkerTaskLen  int //一个Worker的任务队列的长度
}

var GlobalObject *GlobalObj

func (g *GlobalObj) ReLoad() {
	data, err := ioutil.ReadFile("conf/conf.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &GlobalObject)
	if err != nil {
		panic(err)
	}
}
