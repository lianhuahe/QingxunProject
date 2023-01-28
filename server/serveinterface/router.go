package serveinterface

type IRouter interface {
	//将要返回给客户端的处理结果发送到msgbuf_chan
	PreHandleWithMsg(msg *[]byte, msgbuf_chan chan []byte) //业务处理之前
	PreHandle(msgbuf_chan chan []byte)
	HandleWithMsg(msg *[]byte, msgbuf_chan chan []byte) //业务处理
	Handle(msgbuf_chan chan []byte)
	PostHandleWithMsg(msg *[]byte, msgbuf_chan chan []byte) //业务处理之后
	PostHandle(msgbuf_chan chan []byte)
}
