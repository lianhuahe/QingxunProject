package serve

import (
	//protologin "Sample_Douyin_Server/protobuf/user/login"
	serveinterface "Sample_Douyin_Server/server/serveinterface"
	globalobj "Sample_Douyin_Server/server/utils"
	"encoding/json"
	"fmt"
	"net"
	"strconv"
	//"github.com/golang/protobuf/proto"
)

type Connection struct {
	conn_server serveinterface.IServer //该连接所属的server
	conn        *net.TCPConn           //该连接的套接字
	conn_id     uint32                 //连接的身份id
	msgbuf_chan chan []byte            //与写协程间的管道
	msghandle   serveinterface.IMSGHandle
	isclosed    bool //连接状态
}

func (c *Connection) StartReader() {
	fmt.Println("connection", c.conn_id, " start read")
	//读业务结束说明该连接可以断开了
	defer c.Stop()

	for {
		buf := make([]byte, 512)
		n, err := c.conn.Read(buf)
		if err != nil {
			continue
		}
		fmt.Println("connection", c.conn_id, " recv:\n ", string(buf[:n]))
		//TODO拆包解包读数据
		//测试能不能登录
		smsg := "ok"
		var uid int64 = 1
		rsp := LoginResponse{
			StatusCode: 0,
			StatusMsg:  &smsg,
			Token:      &smsg,
			UserID:     &uid,
		}
		data, _ := json.Marshal(rsp)
		lenth := strconv.Itoa(len(data))
		response := "HTTP/1.1 200 OK\r\nContent-Length: " + lenth + "\r\nConnection: keep-alive\r\n\r\n" + string(data)
		fmt.Println(response)
		c.conn.Write([]byte(response))
		/*cont := make([]byte, 512)
		var code uint32 = 1
		req := CreateRequest(0, cont, code, c)
		c.msghandle.SendMSGToHandle(c.conn_id, req)*/
	}
}

func (c *Connection) Start() {
	fmt.Println("connection", c.conn_id, " start")
	//启动读业务
	c.StartReader()
	//启动写业务
}

func (c *Connection) Stop() {
	if c.isclosed {
		return
	}
	fmt.Println("connection", c.conn_id, " closed")
	c.isclosed = true
	c.conn.Close()
	close(c.msgbuf_chan)
}

func (c *Connection) GetTCPConn() *net.TCPConn {
	return c.conn
}

func (c *Connection) GetConnId() uint32 {
	return c.conn_id
}

func (c *Connection) GetBufChan() chan []byte {
	return c.msgbuf_chan
}

func CreateConnection(server serveinterface.IServer, conn *net.TCPConn, id uint32) serveinterface.IConnection {
	c := &Connection{
		conn_server: server,
		conn:        conn,
		conn_id:     id,
		msgbuf_chan: make(chan []byte, globalobj.GlobalObject.MsgChanLen),
		isclosed:    true,
	}
	return c
}
