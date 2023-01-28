package serve

import (
	serveinterface "Sample_Douyin_Server/server/serveinterface"
	globalobj "Sample_Douyin_Server/server/utils"
	"errors"
	"fmt"
	"net"
)

type Server struct {
	ip_version string
	ip         string
	port       int
	conn_count int
	exit_chan  chan struct{}
	MSG_handle serveinterface.IMSGHandle
}

func (s *Server) Start() {
	fmt.Printf("[START] listenner at IP: %s, Port %d is starting\n", s.ip, s.port)
	s.exit_chan = make(chan struct{})
	//开启协程进行监听
	go func() {
		//TODO数据库连接池初始化
		//TODO部分数据缓存
		server_addr, err := net.ResolveTCPAddr(s.ip_version, fmt.Sprintf("%s:%d", s.ip, s.port))
		if err != nil {
			fmt.Println("resolve tcp addr err: ", err)
			return
		}
		listener, err := net.ListenTCP(s.ip_version, server_addr)
		if err != nil {
			panic(err)
		}
		fmt.Println("listening...")
		go func() {
			var cid uint32 = 1
			for {
				//超过最大连接数
				if s.conn_count >= globalobj.GlobalObject.MaxConn {
					continue
				}
				client_conn, err := listener.AcceptTCP()
				if err != nil {
					if errors.Is(err, net.ErrClosed) {
						fmt.Println("Listener closed")
						return
					}
					fmt.Println("Accept Error")
					continue
				}
				deal_conn := CreateConnection(s, client_conn, cid)
				cid++
				go deal_conn.Start()
			}
		}()
		select {
		case <-s.exit_chan:
			err := listener.Close()
			if err != nil {
				fmt.Println("Listener close err ", err)
			}
		}
	}()
}

func (s *Server) Serve() {
	s.Start()
	//TODO
	//阻塞，防止结束后导致Start()也退出
	select {}
}

func (s *Server) Stop() {
	fmt.Printf("[STOP]")
	s.exit_chan <- struct{}{}
	close(s.exit_chan)
}

func CreateServer() serveinterface.IServer {
	globalobj.GlobalObject.ReLoad()
	s := &Server{
		ip_version: "tcp4",
		ip:         globalobj.GlobalObject.Host,
		port:       globalobj.GlobalObject.TcpPort,
		exit_chan:  nil,
		MSG_handle: CreateMSGHandle(),
	}
	return s
}
