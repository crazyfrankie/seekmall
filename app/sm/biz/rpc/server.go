package rpc

import (
	"github.com/crazyfrankie/seekmall/app/sm/biz/service"
	"github.com/crazyfrankie/seekmall/app/sm/config"
	"google.golang.org/grpc"
	"net"
)

type Server struct {
	*grpc.Server
	Addr string
}

func NewServer(sm *service.SmsServer) *Server {
	s := grpc.NewServer()
	sm.RegisterServer(s)

	return &Server{
		Server: s,
		Addr:   config.GetConf().Server.Addr,
	}
}

func (s *Server) Serve() error {
	conn, err := net.Listen("tcp", s.Addr)
	if err != nil {
		return err
	}

	return s.Server.Serve(conn)
}
