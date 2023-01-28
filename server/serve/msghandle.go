package serve

import (
	serveinterface "Sample_Douyin_Server/server/serveinterface"
)

var APIs map[uint32]serveinterface.IRouter //处理方法的路由集合

func init() {
	APIs = make(map[uint32]serveinterface.IRouter)
}

// 用于在完成一个业务操作后添加路由方法
func AddRouter(route_code uint32, route serveinterface.IRouter) {
	APIs[route_code] = route
}

/*
//下面是实现一个Router的方法
type Router struct {
	BaseRouter
}

func ADD()  {
	APIs.AddRouter(1,&Router{

	})
}
//按需重写BaseRouter的函数，
*/

type MSGHandle struct {
	workerpool *WorkerPool //所拥有的工作池
}

func CreateMSGHandle() *MSGHandle {
	return &MSGHandle{
		workerpool: CreateWorkerPool(),
	}
}

func (m *MSGHandle) SendMSGToHandle(mid uint32, request serveinterface.IRequest) {
	//TODO 负载均衡算法，目前先用一个简单的hash，之后再写
	loc := mid % uint32(m.workerpool.workerpool_size)
	m.workerpool.task_queue[loc] <- request
}
