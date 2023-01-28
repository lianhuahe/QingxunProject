package serveinterface

import "net"

type IConnection interface {
	Start()                   //启动连接
	Stop()                    //关闭连接
	GetTCPConn() *net.TCPConn //获取该连接的套接字
	GetConnId() uint32        //获取连接id
	GetBufChan() chan []byte  //获取与写协程间的管道
}
