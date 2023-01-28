package serveinterface

type IMSGHandle interface {
	SendMSGToHandle(mid uint32, request IRequest) //将消息发送给MSGHandle来处理
}
