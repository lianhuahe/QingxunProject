package serve

//serveinterface "Sample_Douyin_Server/server/serveinterface"

type BaseRouter struct{}

func (r *BaseRouter) PreHandleWithMsg(msg *[]byte, msgbuf_chan chan []byte) {}

func (r *BaseRouter) PreHandle(msgbuf_chan chan []byte) {}

func (r *BaseRouter) HandleWithMsg(msg *[]byte, msgbuf_chan chan []byte) {}

func (r *BaseRouter) Handle(msgbuf_chan chan []byte) {}

func (r *BaseRouter) PostHandleWithMsg(msg *[]byte, msgbuf_chan chan []byte) {}

func (r *BaseRouter) PostHandle(msgbuf_chan chan []byte) {}
