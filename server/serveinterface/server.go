package serveinterface

type IServer interface {
	Start() //启动服务器
	Serve() //开启服务
	Stop()  //停止服务器
}
