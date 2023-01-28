package serve

import (
	//serveinterface "Sample_Douyin_Server/server/serveinterface"
	"Sample_Douyin_Server/server/serveinterface"
	globalobj "Sample_Douyin_Server/server/utils"
	"fmt"
)

type WorkerPool struct {
	task_queue      []chan serveinterface.IRequest
	workerpool_size int
}

func StartWork(task_queue chan serveinterface.IRequest) {
	for {
		select {
		case request := <-task_queue:
			request.DoRequest()
		}
	}
}

func (w *WorkerPool) Start() {
	for i := 0; i < w.workerpool_size; i++ {
		fmt.Println("Worker", i, "is created")

		go StartWork(w.task_queue[i])
	}
}

func (w *WorkerPool) Stop() {

}

func CreateWorkerPool() *WorkerPool {
	return &WorkerPool{
		workerpool_size: globalobj.GlobalObject.WorkerPoolSize,
		task_queue:      make([]chan serveinterface.IRequest, globalobj.GlobalObject.WorkerTaskLen),
	}
}
