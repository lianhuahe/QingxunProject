package serve

import (
	serveinterface "Sample_Douyin_Server/server/serveinterface"
)

type Request struct {
	conn    serveinterface.IConnection //所属连接
	msg_len int                        //消息长度
	msg     []byte                     //消息内容
	API     serveinterface.IRouter     //处理该消息的api
}

func (r *Request) DoRequest() {
	r.API.PreHandle(r.conn.GetBufChan())
	r.API.PreHandleWithMsg(&r.msg, r.conn.GetBufChan())
	r.API.Handle(r.conn.GetBufChan())
	r.API.HandleWithMsg(&r.msg, r.conn.GetBufChan())
	r.API.PostHandle(r.conn.GetBufChan())
	r.API.PostHandleWithMsg(&r.msg, r.conn.GetBufChan())
}

func CreateRequest(msg_len int, msg []byte, API_code uint32, conn serveinterface.IConnection) serveinterface.IRequest {
	return &Request{
		conn:    conn,
		msg_len: msg_len,
		msg:     msg,
		API:     APIs[API_code],
	}
}
